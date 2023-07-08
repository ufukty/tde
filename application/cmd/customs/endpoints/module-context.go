package endpoints

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"tde/config"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var (
	ErrRequiredRequestPatameterNotFound = errors.New("required request parameter (either in url, header or body)")
)

type ContextRequest struct {
	ArchiveId string
	Folder    string
	File      string
	Function  string
}

type ContextResponse struct {
	Context *context.Context `json:"context"`
}

func (req *ContextRequest) NewRequest() (*http.Request, error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed on serialization")
	}
	httpRequest, err := http.NewRequest(
		config.CustomsModuleContext.Method.String(),
		config.CustomsModuleContext.Url(),
		buffer,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed on creating request object")
	}
	return httpRequest, nil
}

func BindRequest(r *http.Request) (bind *ContextRequest, err error) {
	var (
		vars map[string]string
		ok   bool
	)

	vars = mux.Vars(r)

	if bind.ArchiveId, ok = vars["id"]; !ok {
		return nil, errors.Wrap(ErrRequiredRequestPatameterNotFound, "id")
	}
	if bind.Folder, ok = vars["package"]; !ok {
		return nil, errors.Wrap(ErrRequiredRequestPatameterNotFound, "package")
	}
	if bind.File, ok = vars["file"]; !ok {
		return nil, errors.Wrap(ErrRequiredRequestPatameterNotFound, "file")
	}
	if bind.Function, ok = vars["function"]; !ok {
		return nil, errors.Wrap(ErrRequiredRequestPatameterNotFound, "function")
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

func (res *ContextResponse) SerializeIntoResponseWriter(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		return errors.Wrap(err, "failed on serialization")
	}
	return nil
}

func (res *ContextResponse) DeserializeResponse(r *http.Response) error {
	err := json.NewDecoder(r.Body).Decode(res)
	if err != nil {
		return errors.Wrap(err, "failed on parsing the response body")
	}
	return nil
}

// func SendContext(req ContextRequest) (res ContextResponse, err error) {
// 	var r *http.Request
// 	r, err = req.NewRequest()
// 	return res, nil
// }

func (em EndpointsManager) HandleContext(w http.ResponseWriter, r *http.Request) {
}
