package endpoints

import (
	"context"
	"net/http"
	"net/url"
	"tde/config"
	"tde/i18n"
	"tde/internal/microservices/utilities"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type ContextRequest struct {
	ArchiveId string `url:"archive_id"`
	Folder    string `url:"package"`
	File      string `url:"file"`
	Function  string `url:"function"`
}

type ContextResponse struct {
	Context *context.Context `json:"context"`
}

func BindRequest(r *http.Request) (bind *ContextRequest, err error) {
	var (
		vars map[string]string
		ok   bool
	)

	vars = mux.Vars(r)

	if bind.ArchiveId, ok = vars["id"]; !ok {
		return nil, errors.Wrap(i18n.ErrMissingParameter, "id")
	}
	if bind.Folder, ok = vars["package"]; !ok {
		return nil, errors.Wrap(i18n.ErrMissingParameter, "package")
	}
	if bind.File, ok = vars["file"]; !ok {
		return nil, errors.Wrap(i18n.ErrMissingParameter, "file")
	}
	if bind.Function, ok = vars["function"]; !ok {
		return nil, errors.Wrap(i18n.ErrMissingParameter, "function")
	}

	if typedArchiveId, err := uuid.Parse(bind.ArchiveId); err != nil {
		return nil, errors.Wrapf(err, "checking sent ArchiveId with uuid package (%s)", bind.ArchiveId)
	} else {
		bind.ArchiveId = typedArchiveId.String()
	}
	if bind.Folder, err = url.PathUnescape(bind.Folder); err != nil {
		return nil, errors.Wrapf(err, "decoding urlencode for Folder (%s)", bind.Folder)
	}
	if bind.File, err = url.PathUnescape(bind.File); err != nil {
		return nil, errors.Wrapf(err, "decoding urlencode for File (%s)", bind.File)
	}
	if bind.Function, err = url.PathUnescape(bind.Function); err != nil {
		return nil, errors.Wrapf(err, "decoding urlencode for Function (%s)", bind.Function)
	}

	return bind, nil
}

func ContextSend(bq *ContextRequest) (*ContextResponse, error) {
	return utilities.Send[ContextRequest, ContextResponse](config.CustomsModuleContext, bq)
}

func (em EndpointsManager) ContextHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
