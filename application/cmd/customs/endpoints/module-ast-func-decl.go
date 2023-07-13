package endpoints

import (
	"os"
	"tde/config"
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
)

type AstFuncDeclRequest struct {
	ArchiveId string `url:"id"`
	Package   string `url:"package"`
	File      string `url:"file"`
	Function  string `url:"function"`
}

type AstFuncDeclResponse struct {
	FuncDecl *ast.FuncDecl `json:"function"`
}

func AstFuncDeclSend(q AstFuncDeclRequest) (*AstFuncDeclResponse, error) {
	return utilities.Send[AstFuncDeclRequest, AstFuncDeclResponse](config.CustomsModuleAstFuncDecl, q)
}

func (em EndpointsManager) AstFuncDeclHandler() func(w http.ResponseWriter, r *http.Request) {
	var (
		ErrAstConversionFailed = errors.New("AST convertion has failed")
		ErrFuncDeclNotFound    = errors.New("Requested function declaration is not found in the file")
	)

	var transferReadyAstOfFuncDecl = func(bq *AstFuncDeclRequest) (*ast.FuncDecl, error) {
		var (
			file *ast.File
			err  error
			wd   string
		)
		wd, err = os.Getwd()
		if err != nil {
			return nil, errors.Wrap(err, "Can't get the working directory")
		}
		err = os.Chdir(bq.Package)
		if err != nil {
			return nil, errors.Wrap(err, "Can't switch to the directory of requested package")
		}
		defer os.Chdir(wd)

		file, err = parser.ParseFile(token.NewFileSet(), bq.File, nil, parser.AllErrors)
		if err != nil {
			return nil, errors.Wrap(ErrAstConversionFailed, err.Error())
		}

		var funcDecl *ast.FuncDecl
		ast.Inspect(file, func(n ast.Node) bool {
			if n, ok := n.(*ast.FuncDecl); ok {
				if n.Name.Name == bq.Function {
					funcDecl = n
				}
			}
			return funcDecl == nil
		})
		if funcDecl == nil {
			return nil, ErrFuncDeclNotFound
		}
		return clean.FuncDecl(funcDecl), nil
	}

	var (
		ErrUrlDecodingPackage      = errors.New("Failed at decoding requested package path")
		ErrUrlDecodingFilename     = errors.New("Failed at decoding requested filename")
		ErrUrlDecodingFunctionName = errors.New("Failed at decoding requested function name")
		ErrEvilPathFound           = errors.New("Links are not allowed at paths")
	)

	var sanitizePaths = func(bq *AstFuncDeclRequest) error {
		var err error

		if bq.Package, err = url.PathUnescape(bq.Package); err != nil {
			return ErrUrlDecodingPackage
		}

		if bq.File, err = url.PathUnescape(bq.File); err != nil {
			return ErrUrlDecodingFilename
		}

		if bq.Function, err = url.PathUnescape(bq.Function); err != nil {
			return ErrUrlDecodingFunctionName
		}

		bq.Package = filepath.Clean(bq.Package)
		bq.File = filepath.Clean(filepath.Base(bq.File))

		if utilities.IsEvilPath(bq.Package) {
			return errors.Wrap(ErrEvilPathFound, bq.Package)
		}
		if utilities.IsEvilPath(bq.File) {
			return errors.Wrap(ErrEvilPathFound, bq.File)
		}

		var _, _, extract = em.vm.FindPath(bq.ArchiveId)
		bq.Package = filepath.Join(extract, bq.Package)

		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err      error
			bq       *AstFuncDeclRequest
			funcDecl *ast.FuncDecl
		)

		if bq, err = utilities.ParseRequest[AstFuncDeclRequest](r); err != nil {
			log.Println(errors.Wrap(err, ResponseMalformedRequest))
			http.Error(w, ResponseMalformedRequest, http.StatusBadRequest)
			return
		}

		if bundle, zip, extract := em.vm.CheckIfExists(bq.ArchiveId); !(bundle && zip && extract) {
			log.Println("Failed at check the module if exists")
			http.NotFound(w, r)
			return
		}

		if err := sanitizePaths(bq); err != nil {
			log.Println(errors.Wrap(err, "sanitizing parameters"))
			switch {
			case errors.Is(err, ErrUrlDecodingPackage),
				errors.Is(err, ErrUrlDecodingFilename),
				errors.Is(err, ErrUrlDecodingFunctionName):
				http.Error(w, "Woops. Check your parameters or try again later", http.StatusInternalServerError)
			case errors.Is(err, ErrEvilPathFound):
				http.Error(w, errors.Cause(err).Error(), http.StatusBadRequest)
			default:
				http.Error(w, "Woops. Try again later.", http.StatusInternalServerError)
			}
			return
		}

		if funcDecl, err = transferReadyAstOfFuncDecl(bq); err != nil {
			log.Println(errors.Wrap(err, "requesting ast representation of package at path"))
			switch {
			case errors.Is(err, ErrAstConversionFailed):
				http.Error(w, errors.Cause(err).Error(), http.StatusInternalServerError)

			case errors.Is(err, ErrFuncDeclNotFound):
				http.Error(w, errors.Cause(err).Error(), http.StatusBadRequest)

			default:
				http.Error(w, "Woops.", http.StatusInternalServerError)
			}
			return
		}

		var bs = &AstFuncDeclResponse{
			FuncDecl: funcDecl,
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
