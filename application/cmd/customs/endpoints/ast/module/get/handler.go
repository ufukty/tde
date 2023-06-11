package ast

import (
	"fmt"
	"net/http"
	volume_manager "tde/cmd/customs/internal/volume-manager"
	"tde/internal/microservices/errors/bucket"
	"tde/internal/microservices/errors/detailed"
	"tde/internal/microservices/logger"
)

var (
	volumeManager *volume_manager.VolumeManager
	log           = logger.NewLogger("customs/endpoints/module/ast/get")
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
	if de := assertHeader(r, "Content-Type", "application/json"); de != nil {
		errs.Add(de)
	}
	return errs
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		bucket  *bucket.Bucket
		bindReq = new(Request)
	)

	if bucket = checkHeaders(r); bucket.IsAny() {
		var tagged = bucket.Tag(detailed.New("Invalid request", "@Handler"))
		log.Println(tagged.Log())
		http.Error(w, tagged.Error(), http.StatusBadRequest)
		return
	}

	if err = bindReq.ParseRequest(r); err != nil {
		var message = "Request is malformed"
		log.Println(detailed.AddBase(err, message).Log())
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	if ok := volumeManager.CheckIfExists(bindReq.ArchiveId); !ok {
		
	}

}
