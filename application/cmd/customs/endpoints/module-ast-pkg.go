package endpoints

import (
	"fmt"
	"os"
	"tde/config"
	"tde/i18n"
	"tde/internal/astw/clone/clean"
	"tde/internal/microservices/utilities"

	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

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

	var transferReadyAstOfPackage = func(path string) (*ast.Package, error) {
		var (
			pkgs map[string]*ast.Package
			err  error
			wd   string
		)
		wd, err = os.Getwd()
		if err != nil {
			return nil, errors.Wrap(err, "Can't get the working directory")
		}
		err = os.Chdir(path)
		if err != nil {
			return nil, errors.Wrap(err, "Can't switch to the directory of requested package")
		}
		defer os.Chdir(wd)

		pkgs, err = parser.ParseDir(token.NewFileSet(), ".", nil, parser.AllErrors)
		if err != nil {
			return nil, errors.Wrap(i18n.ErrAstConversionFailed, err.Error())
		}
		if l := len(maps.Keys(pkgs)); l == 0 {
			return nil, i18n.ErrNoPackagesFound
		} else if l > 1 {
			return nil, i18n.ErrMultiplePackagesFound
		}

		return clean.Package(pkgs[maps.Keys(pkgs)[0]]), nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err error
			bq  *AstPackageRequest
			pkg *ast.Package
		)

		if bq, err = utilities.ParseRequest[AstPackageRequest](r); err != nil {
			log.Println(errors.Wrap(err, i18n.MalformedRequest))
			http.Error(w, i18n.MalformedRequest, http.StatusBadRequest)
			return
		}

		if bundle, zip, extract := em.vm.CheckIfExists(bq.ArchiveId); !(bundle && zip && extract) {
			log.Println("Failed at check if exists")
			http.NotFound(w, r)
			return
		}

		bq.Package, err = url.PathUnescape(bq.Package)
		if err != nil {
			var response = "Failed at decoding requested package path"
			errors.Wrap(err, response)
			http.Error(w, response, http.StatusBadRequest)
		}

		var folderpath = filepath.Clean(bq.Package)
		if utilities.IsEvilPath(folderpath) {
			var response = fmt.Sprintf("Links are not allowed at paths: '%s'\n", folderpath)
			fmt.Println(response)
			http.Error(w, response, http.StatusBadRequest)
		}

		var _, _, extract = em.vm.FindPath(bq.ArchiveId)
		folderpath = filepath.Clean(filepath.Join(extract, bq.Package))

		if pkg, err = transferReadyAstOfPackage(folderpath); err != nil {
			log.Println(errors.Wrap(err, "requesting ast representation of package at path"))
			switch {
			case errors.Is(err, i18n.ErrAstConversionFailed):
				http.Error(w, errors.Cause(err).Error(), http.StatusInternalServerError)

			case errors.Is(err, i18n.ErrMultiplePackagesFound):
				http.Error(w, errors.Cause(err).Error(), http.StatusBadRequest)

			case errors.Is(err, i18n.ErrNoPackagesFound):
				http.Error(w, errors.Cause(err).Error(), http.StatusNotFound)

			default:
				http.Error(w, "Woops.", http.StatusInternalServerError)
			}
			return
		}

		var bs = &AstPackageResponse{
			Package: pkg,
		}

		if err = utilities.WriteJsonResponse(bs, w); err != nil {
			var response = "Could not write the request body"
			log.Println(errors.Wrap(err, response))
			http.Error(w, response, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
