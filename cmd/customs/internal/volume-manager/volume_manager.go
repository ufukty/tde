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

// returns eg. 65/36/f1/24/b8/56/5a/ad/8c/cc/22/ea/c3/7d/8e/63/filename
func (vm *VolumeManager) CreateDestPath(filename string) (path string, err error) {
	slicedSubdirs := strings.Join(utilities.StringFold(strings.ReplaceAll(filename, "-", ""), 2), "/")
	dirPath := filepath.Join(vm.root, slicedSubdirs)
	err = os.MkdirAll(dirPath, 0700)
	if err != nil {
		return "", errors.Wrap(err, "Could not create folders in created destination path")
	}
	path = filepath.Join(dirPath, filename)
	return
}
