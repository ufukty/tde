package endpoints

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"tde/internal/microservices/errors/detailed"
	"tde/internal/microservices/utilities"

	"golang.org/x/exp/maps"
)

//go:generate serdeser module-ast-pkg.go
type AstPackageRequest struct {
	ArchiveId string `json:"archive_id"`
	Folder    string `json:"folder"`
}

type AstPackageResponse struct {
	Package *ast.Package `json:"package"`
}

func (em EndpointsManager) AstPackageHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err error
			// bucket  *bucket.Bucket
			bindReq = new(AstPackageRequest)
			bindRes = new(AstPackageResponse)
			pkgs    map[string]*ast.Package
		)

		// if bucket = utilities.CheckHeaders(r); bucket.IsAny() {
		// 	var tagged = bucket.Tag(detailed.New("Invalid request", "@Handler"))
		// 	log.Println(tagged.Log())
		// 	http.Error(w, tagged.Error(), http.StatusBadRequest)
		// 	return
		// }

		// if err = bindReq.ParseRequest(r); err != nil {
		// 	var message = "Request is malformed"
		// 	log.Println(detailed.AddBase(err, message).Log())
		// 	http.Error(w, message, http.StatusBadRequest)
		// 	return
		// }

		if bundle, zip, extract := em.vm.CheckIfExists(bindReq.ArchiveId); !(bundle && zip && extract) {
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
}
