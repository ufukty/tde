package endpoints

import (
	"tde/config"
	"tde/internal/microservices/errors/detailed"
	"tde/internal/microservices/utilities"

	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"net/http"
	"path/filepath"

	"golang.org/x/exp/maps"
)

//go:generate serdeser module-ast-pkg.go
type AstPackageRequest struct {
	ArchiveId string `url:"id"`
	Package   string `url:"package"`
}

type AstPackageResponse struct {
	Package *ast.Package `json:"package"`
}

func AstPackageSend(q AstPackageRequest) (*AstPackageResponse, error) {
	return utilities.Send[AstPackageRequest, AstPackageResponse](config.CustomsModuleAstPackage, q)
}

func (em EndpointsManager) AstPackageHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err error
			bq  *AstPackageRequest
		)

		if bq, err = utilities.ParseRequest[AstPackageRequest](r); err != nil {
			var message = "Request is malformed"
			log.Println(detailed.AddBase(err, message).Log())
			http.Error(w, message, http.StatusBadRequest)
			return
		}

		if bundle, zip, extract := em.vm.CheckIfExists(bq.ArchiveId); !(bundle && zip && extract) {
			log.Println(detailed.New("Not found", "volumeManager.CheckIfExists").Log())
			http.NotFound(w, r)
			return
		}

		var folderpath = filepath.Clean(bq.Package)
		if utilities.IsEvilPath(folderpath) {
			log.Printf("requested path: '%s'\n", folderpath)
			http.Error(w, "Bad request. Bad. Bad.", http.StatusBadRequest)
		}

		var pkgs map[string]*ast.Package
		pkgs, err = parser.ParseDir(token.NewFileSet(), folderpath, nil, parser.AllErrors)
		if err != nil {
			http.Error(w, "AST convertion has failed", http.StatusInternalServerError)
			return
		}

		if l := len(maps.Keys(pkgs)); l == 0 {
			http.NotFound(w, r)
		} else if l > 1 {
			log.Printf("more than 1 package found in '%s' => '%s'\n", bq.ArchiveId, bq.Package)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		var bs = &AstPackageResponse{
			Package: pkgs[maps.Keys(pkgs)[0]],
		}

		w.WriteHeader(http.StatusOK)
		if err = utilities.WriteJsonResponse(bs, w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	}
}
