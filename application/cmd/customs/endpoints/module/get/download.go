package module

import (
	"fmt"
	"tde/cmd/customs/internal/utilities"
	volume_manager "tde/cmd/customs/internal/volume-manager"
	"tde/internal/microservices/errors/bucket"
	"tde/internal/microservices/errors/detailed"
	"tde/internal/microservices/logger"
	"time"

	"net/http"
	"os"

	"github.com/pkg/errors"
)

const (
	HTTPErrBadRequest = "Bad request"
)

var (
	volumeManager *volume_manager.VolumeManager
	log           = logger.NewLogger("customs/endpoints/module/get/handler")
)

func RegisterVolumeManager(vm *volume_manager.VolumeManager) {
	volumeManager = vm
}

func assertHeader(r *http.Request, headerField, want string) *detailed.DetailedError {
	var got = r.Header.Get(headerField)
	var expected = want
	if expected != got {
		return detailed.New(
			fmt.Sprintf("Value for %s header field is unexpected.", headerField),
			fmt.Sprintf("Expected %s, Got %s", expected, got),
		)
	}
	return nil
}

func checkHeaders(r *http.Request) *bucket.Bucket {
	var errs = new(bucket.Bucket)
	if de := assertHeader(r, "Content-Type", "application/json; charset=utf-8"); de != nil {
		errs.Add(de)
	}
	return errs
}

func Handler(w http.ResponseWriter, r *http.Request) {

	var (
		parameters  = new(Request)
		err         error
		fileHandler *os.File
		digest      string
	)

	if bucket := checkHeaders(r); bucket.IsAny() {
		var tagged = bucket.Tag(detailed.New("Invalid request", "@Handler"))
		log.Println(tagged.Log())
		http.Error(w, tagged.Error(), http.StatusBadRequest)
		return
	}

	if err = parameters.ParseRequest(r); err != nil {
		var message = "Invalid request body"
		log.Println(errors.Wrap(err, message))
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	if ok := volumeManager.CheckIfExists(parameters.ArchiveId); !ok {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	var filePath = volumeManager.FindPath(parameters.ArchiveId)
	fileHandler, err = os.Open(filePath)
	if err != nil {
		var derr = detailed.New("File not found", err.Error())
		log.Println(derr.Log())
		http.Error(w, derr.Error(), http.StatusNotFound)
		return
	}
	defer fileHandler.Close()

	http.ServeContent(w, r, "file.zip", time.Now(), fileHandler)

	digest, err = utilities.MD5(fileHandler)
	if err != nil {
		http.Error(w, "Could not calculate the checksum of file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Digest", fmt.Sprintf("md5=%s", digest))
	w.Header().Set("Content-Disposition", "attachment; filename=file.zip")

}
