package endpoints

import (
	"path/filepath"
	"tde/config"
	"tde/i18n"
	astutils "tde/internal/astw/astwutl"
	"tde/internal/astw/clone/clean"
	"tde/internal/astw/traverse"
	"tde/internal/evolution/genetics/mutation/cfg/ctxres"
	"tde/internal/evolution/genetics/mutation/cfg/ctxres/context"
	"tde/internal/microservices/utilities"

	"go/ast"
	"log"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ContextRequest struct {
	ArchiveId string `url:"id"`
	Package   string `url:"package"`
	File      string `url:"file"`
	Function  string `url:"function"`
}

type ContextResponse struct {
	Context *context.Context `json:"context"`
}

func ContextSend(bq *ContextRequest) (*ContextResponse, error) {
	return utilities.Send[ContextRequest, ContextResponse](config.CustomsModuleContext, bq)
}

func (em EndpointsManager) ContextHandler() func(w http.ResponseWriter, r *http.Request) {
	var (
		ErrUrlDecodingPackage      = errors.New("Failed at decoding requested package path")
		ErrUrlDecodingFilename     = errors.New("Failed at decoding requested filename")
		ErrUrlDecodingFunctionName = errors.New("Failed at decoding requested function name")
		ErrEvilPathFound           = errors.New("Links are not allowed at paths")
	)

	var sanitizeRequest = func(bq *ContextRequest) error {
		var err error

		if typedArchiveId, err := uuid.Parse(bq.ArchiveId); err != nil {
			return errors.Wrap(i18n.ErrInputSanitization, errors.Wrap(err, "ArchiveId").Error())
		} else {
			bq.ArchiveId = typedArchiveId.String()
		}

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
			err error
			bq  *ContextRequest
		)

		if bq, err = utilities.ParseRequest[ContextRequest](r); err != nil {
			log.Println(errors.Wrap(err, i18n.MalformedRequest))
			http.Error(w, "Woops.", http.StatusBadRequest)
			return
		}

		if err = sanitizeRequest(bq); err != nil {
			log.Println(errors.Wrap(err, "sanitizating the request"))
			switch {
			case errors.Is(err, i18n.ErrInputSanitization):
				http.Error(w, "Woops.", http.StatusBadRequest)
			}
			return
		}

		var (
			pkg      *ast.Package
			file     *ast.File
			funcDecl *ast.FuncDecl
			ctx      *context.Context
			ok       bool
		)

		if pkg, err = astutils.LoadPackageFromDir(bq.Package); err != nil {
			log.Println(errors.Wrap(i18n.ErrAstConversionFailed, err.Error()))
			http.Error(w, "Woops.", http.StatusBadRequest)
			return
		}

		pkg = clean.Package(pkg)
		if file, ok = pkg.Files[bq.File]; !ok {
			log.Println(i18n.ErrFileNotFoundInPackage)
			http.Error(w, "Woops.", http.StatusBadRequest)
			return
		}

		funcDecl, err = astutils.FindFuncDecl(file, bq.Function)
		if err != nil {
			log.Println(errors.Wrap(i18n.ErrAstConversionFailed, err.Error()))
			http.Error(w, "Woops.", http.StatusBadRequest)
			return
		}

		var tFuncDecl = traverse.GetTraversableNodeForASTNode(funcDecl)
		var funcBody = traverse.GetTraversableNodeForASTNode(funcDecl.Body).GetTraversableSubnodes()
		var funcBodyLastLine = funcBody[len(funcBody)-1]

		ctx, err = ctxres.GetContextForSpot(pkg, tFuncDecl, funcBodyLastLine)
		if err != nil {
			log.Println(errors.Wrap(errors.New("Could not get context for choosen spot"), err.Error()))
			http.Error(w, "Woops.", http.StatusBadRequest)
			return
		}

		var bs = ContextResponse{
			Context: ctx,
		}

		err = utilities.WriteJsonResponse(bs, w)
	}
}
