package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"mime"
	"net/http"
	"path/filepath"
	volume_manager "tde/cmd/customs/internal/volume-manager"
	"tde/internal/microservices/errors/bucket"
	"tde/internal/microservices/errors/detailed"
	"tde/internal/microservices/logger"
	"tde/internal/microservices/utilities"

	"golang.org/x/exp/maps"
)

var (
	vm  *volume_manager.VolumeManager
	log = logger.NewLogger("customs/endpoints/module/ast/get")
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
		bindRes = new(Response)
		pkgs    map[string]*ast.Package
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

	if bundle, zip, extract := vm.CheckIfExists(bindReq.ArchiveId); !(bundle && zip && extract) {
		log.Println(detailed.New("Not found", "volumeManager.CheckIfExists").Log())
		http.NotFound(w, r)
		return
	}

	var folderpath = filepath.Clean(bindReq.Folder)
	if utilities.IsEvilPath(folderpath) {
		log.Printf("requested path: '%s'\n", folderpath)
		http.Error(w, "Bad request. Bad. Bad.", http.StatusBadRequest)
	}

	pkgs, err = parser.ParseDir(token.NewFileSet(), folderpath, nil, parser.AllErrors)
	if err != nil {
		http.Error(w, "AST convertion has failed", http.StatusInternalServerError)
		return
	}

	if l := len(maps.Keys(pkgs)); l == 0 {
		http.NotFound(w, r)
	} else if l > 1 {
		log.Printf("more than 1 package found in '%s' => '%s'\n", bindReq.ArchiveId, bindReq.Folder)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	bindRes.Package = pkgs[maps.Keys(pkgs)[0]]

	w.Header().Set("Content-Type", mime.TypeByExtension("json"))
	w.WriteHeader(http.StatusOK)
	if err = bindRes.SerializeIntoResponseWriter(w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
