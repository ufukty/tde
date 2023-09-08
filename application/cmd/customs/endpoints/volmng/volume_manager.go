package volmng

import (
	"tde/internal/evolution/evaluation/archive"
	"tde/internal/utilities"

	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type VolumeManager struct {
	root string
}

func NewVolumeManager(root string) *VolumeManager {
	return &VolumeManager{
		root: root,
	}
}

func (vm *VolumeManager) CreateUniqueFilename() string {
	return uuid.NewString()
}

func pathSlice(uuid string) string {
	return strings.Join(utilities.StringFold(strings.ReplaceAll(uuid, "-", ""), 2), "/")
}

// non-NotExists kind errors resolve to false too
func (vm *VolumeManager) CheckIfExists(archiveID string) (bundle, zip, extract bool) {
	var bundle_, zip_, extract_ = vm.FindPath(archiveID)
	var err error
	_, err = os.Stat(bundle_)
	bundle = err == nil
	_, err = os.Stat(zip_)
	zip = err == nil
	_, err = os.Stat(extract_)
	extract = err == nil
	return
}

// example output:
// bundle : 65/36/../8e/63
// zip    : 65/36/../8e/63/original.zip
// extract: 65/36/../8e/63/extract
func (vm *VolumeManager) FindPath(archiveID string) (bundle, zip, extract string) {
	var sliced = pathSlice(archiveID)
	bundle = filepath.Join(vm.root, sliced)
	extract = filepath.Join(bundle, "extract")
	zip = filepath.Join(bundle, "original.zip")
	return
}

func removeArtifacts(bundlepath string) error {
	if err := os.RemoveAll(bundlepath); err != nil {
		return errors.Wrap(err, "calling os.RemoveAll")
	}
	return nil
}

func (vm *VolumeManager) New(uploadFileHandler io.Reader) (archiveId string, errFile string, err error) {
	archiveId = vm.CreateUniqueFilename()
	var bundle, zip, extract = vm.FindPath(archiveId)

	if err := os.MkdirAll(extract, 0700); err != nil {
		return "", "", errors.Wrap(err, "creating path to put uploaded file")
	}

	out, err := os.Create(zip)
	if err != nil {
		return "", "", errors.Wrap(err, "creating target file for zip")
	}
	var closed = false
	defer func() {
		if !closed {
			out.Close()
		}
	}()

	if _, err = io.Copy(out, uploadFileHandler); err != nil {
		return "", "", errors.Wrap(err, "writing zip file")
	}
	closed = true
	if err = out.Close(); err != nil {
		log.Println("failed closing a file: " + out.Name())
		return "", "", errors.Wrap(err, "closing file")
	}

	if err, errFile := archive.Unarchive(zip, extract); err != nil {
		if errCleaning := removeArtifacts(bundle); errCleaning != nil {
			err = errors.Wrap(err, errors.Wrap(errCleaning, "removing artifacts").Error())
		}
		return "", errFile, err
	}
	return archiveId, "", nil
}
