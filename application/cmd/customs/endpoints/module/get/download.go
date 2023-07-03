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

var (
	vm  *volume_manager.VolumeManager
	log = logger.NewLogger("customs/endpoints/module/get/handler")
)

func RegisterVolumeManager(vm_ *volume_manager.VolumeManager) {
	vm = vm_
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
		err         error
		digest      string
		fileHandler *os.File
		bindReq     = new(Request)
	)

	if bucket := checkHeaders(r); bucket.IsAny() {
		log.Println(errors.Wrap(bucket, "checking hearders"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = bindReq.ParseRequest(r); err != nil {
		log.Println(errors.Wrap(err, "Invalid request body"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var bundleExists, zipExists, extractExists = vm.CheckIfExists(bindReq.ArchiveId)
	if !(bundleExists && zipExists && extractExists) {
		log.Printf("Got asked for non-existent archive '%s'\n", bindReq.ArchiveId)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var _, zip, _ = vm.FindPath(bindReq.ArchiveId)
	fileHandler, err = os.Open(zip)
	if err != nil {
		log.Println(errors.Wrap(err, "opening file to read"))
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	defer fileHandler.Close()

	w.WriteHeader(http.StatusOK)
	http.ServeContent(w, r, "file.zip", time.Now(), fileHandler)

	digest, err = utilities.MD5(fileHandler)
	if err != nil {
		log.Println(errors.Wrap(err, "checking md5sum"))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Digest", fmt.Sprintf("md5=%s", digest))
	w.Header().Set("Content-Disposition", "attachment; filename=file.zip")
}
