package slotmgr

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateSwapFile(path string) (string, error) {
	ifh, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("opening file to read: %w", err)
	}
	defer ifh.Close()
	ofh, err := os.CreateTemp(os.TempDir(), filepath.Base(path)+".*.swp")
	if err != nil {
		return "", fmt.Errorf("creating swap file: %w", err)
	}
	defer ofh.Close()
	_, err = io.Copy(ofh, ifh)
	if err != nil {
		return "", fmt.Errorf("copying contents of original file to swap file: %w", err)
	}
	return ofh.Name(), nil
}

var (
	ErrFromLineIsAfterFileEnds = fmt.Errorf("starting line of replacement section is after document ends")
	ErrToLineIsAfterFileEnds   = fmt.Errorf("ending line of replacement section is after document ends")
)

// replacement of bufio.ScanLines. Difference is that this one returns the "\n" at the end of line
func scanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		return i + 1, data[0 : i+1], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

// replaces lines in range ["from", "to") with "content" of file at "path"
func ReplaceSectionInFile(path string, fromLine, toLine int, content []byte) error {
	dest, err := os.CreateTemp(os.TempDir(), filepath.Base(path)+".*.swp")
	if err != nil {
		return fmt.Errorf("creating empty file in tmp: %w", err)
	}
	defer dest.Close()

	src, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("opening file to read: %w", err)
	}
	defer src.Close()
	srcScanner := bufio.NewScanner(src)
	srcScanner.Split(scanLines)

	for i := 0; i < fromLine; i++ {
		if !srcScanner.Scan() {
			if err := srcScanner.Err(); err != nil {
				return fmt.Errorf("scanning lines in the existing part before replacement section: %w", err)
			}
			return ErrFromLineIsAfterFileEnds
		}
		_, err := dest.Write(srcScanner.Bytes())
		if err != nil {
			return fmt.Errorf("copying existing parts before replacement section starts: %w", err)
		}
	}

	_, err = dest.Write(content)
	if err != nil {
		return fmt.Errorf("writing contents into specified section: %w", err)
	}
	fmt.Fprintln(dest) // re-adding existing new line

	// advancing the scanner to "toLine"
	for i := fromLine; i < toLine; i++ {
		if !srcScanner.Scan() {
			if err := srcScanner.Err(); err != nil {
				return fmt.Errorf("advancing the replacement section: %w", err)
			}
			return ErrToLineIsAfterFileEnds
		}
	}

	for true {
		if !srcScanner.Scan() {
			if err := srcScanner.Err(); err != nil {
				return fmt.Errorf("scanning lines in the existing part after replacement section: %w", err)
			}
			break // EOF, naturally
		}
		_, err := dest.Write(srcScanner.Bytes())
		if err != nil {
			return fmt.Errorf("copying existing parts after replacement section ends: %w", err)
		}
	}

	err = os.Rename(dest.Name(), path)
	if err != nil {
		return fmt.Errorf("moving swap file to exisiting file's location: %w", err)
	}

	return nil
}
