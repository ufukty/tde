package module

import (
	"net/http"
	volume_manager "tde/cmd/customs/internal/volume-manager"

	"testing"
)

func init() {
	vm = volume_manager.NewVolumeManager("test-files/mount")
}

func Test_DownloadHandler(t *testing.T) {
	var bindReq = Request{ArchiveId: "4603c942-5dd3-5912-9874-7df2fb078066"}
	var r, err = bindReq.NewRequest("GET", "/customs/module")
	if err != nil {
		t.Fatalf("prep: %v", err)
	}

	http.DefaultClient.Do(r)
}
