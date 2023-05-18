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

// returns eg. 65/36/f1/24/b8/56/5a/ad/8c/cc/22/ea/c3/7d/8e/63 after creating it as dir
func (vm *VolumeManager) CreateDestPath(archiveID string) (path string, err error) {
	dirPath := filepath.Join(vm.root, pathSlice(archiveID))
	err = os.MkdirAll(dirPath, 0700)
	if err != nil {
		return "", errors.Wrap(err, "Could not create folders in created destination path")
	}
	return dirPath, nil
}
