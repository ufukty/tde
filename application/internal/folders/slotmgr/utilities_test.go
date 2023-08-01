package slotmgr

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/kylelemons/godebug/diff"
	"golang.org/x/exp/slices"
)

func Test_ReplaceSectionInFile(t *testing.T) {
	testcases := []string{
		"testdata/case1",
		"testdata/case2", // ends with a linefeed
		"testdata/case3", // ends with 3 empty lines
	}
	for _, testcase := range testcases {
		t.Run(testcase, func(t *testing.T) {
			inputfile, err := CreateSwapFile(filepath.Join(testcase, "input.txt"))
			if err != nil {
				t.Fatal(fmt.Errorf("%s prep: %w", testcase, err))
			}
			fmt.Printf("swap file for %s: %s\n", testcase, inputfile)
			want, err := os.ReadFile(filepath.Join(testcase, "want.txt"))
			if err != nil {
				t.Fatal(fmt.Errorf("%s prep: %w", testcase, err))
			}
			modifications, err := os.ReadFile(filepath.Join(testcase, "contents.txt"))
			if err != nil {
				t.Fatal(fmt.Errorf("%s prep: %w", testcase, err))
			}
			err = ReplaceSectionInFile(inputfile, 9, 19, modifications) // lines [9, 19)
			if err != nil {
				t.Fatal(fmt.Errorf("%s act: %w", testcase, err))
			}
			got, err := os.ReadFile(inputfile)
			if err != nil {
				t.Fatal(fmt.Errorf("%s assert prep: %w", testcase, err))
			}
			if slices.Compare(want, got) != 0 {
				fmt.Println(testcase, " diff want->got '+', '-', or ' '")
				fmt.Println(diff.Diff(string(want), string(got)))
				t.Error(fmt.Errorf("%s assert", testcase))
			}
		})
	}
}
