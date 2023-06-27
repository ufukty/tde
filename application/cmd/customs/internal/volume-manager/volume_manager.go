package volume_manager

import (
	"tde/internal/utilities"

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

// returns eg. 65/36/f1/24/b8/56/5a/ad/8c/cc/22/ea/c3/7d/8e/63/{original.zip,extract} after creating it as dir
func (vm *VolumeManager) CreateDestPath(archiveID string) (zipPath string, extractPath string, err error) {
	var slicedId = pathSlice(archiveID)
	var bundleDest = filepath.Join(vm.root, slicedId)
	if err = os.MkdirAll(bundleDest, 0700); err != nil {
		return "", "", errors.Wrap(err, "Could not create dir for upload")
	}
	zipPath = filepath.Join(bundleDest, "original.zip")
	extractPath = filepath.Join(bundleDest, "extract")
	if err = os.MkdirAll(extractPath, 0700); err != nil {
		return "", "", errors.Wrap(err, "Could not create dir for extraction")
	}
	return
}

func (vm *VolumeManager) CheckIfExists(archiveID string) bool {
	var (
		path string
		err  error
	)
	path = filepath.Join(vm.root, pathSlice(archiveID), archiveID)
	_, err = os.Stat(path)
	return err != nil // who cares about the condition where the file exists but there is an error
}

func (vm *VolumeManager) FindPath(archiveID string) string {
	return filepath.Join(vm.root, pathSlice(archiveID), archiveID) + ".zip"
}
