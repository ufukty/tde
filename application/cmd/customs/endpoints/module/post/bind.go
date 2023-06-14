package module

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"tde/cmd/customs/internal/utilities"

	"github.com/pkg/errors"
)

//go:generate serdeser bind.go

// NOTE: DON'T use serdeser for multipart/form-data requests
type Request struct {
	body                    *bytes.Buffer
	contentTypeWithBoundary string
}

type Response struct {
	ArchiveID string `json:"archive_id"`
}

func NewRequest(file io.Reader) (*Request, error) {
	var (
		err          error
		digest       string
		fileField    io.Writer
		cheksumField io.Writer
		req          = new(Request)
		mpWriter     *multipart.Writer
	)

	req.body = bytes.NewBuffer([]byte{})
	mpWriter = multipart.NewWriter(req.body)

	var checksumFieldRead = bytes.NewBuffer([]byte{})
	var fileFieldRead = io.TeeReader(file, checksumFieldRead)

	{
		digest, err = utilities.MD5(checksumFieldRead)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to find checksum of file")
		}
		cheksumField, err = mpWriter.CreateFormField("md5sum")
		if err != nil {
			return nil, errors.Wrap(err, "Could not create form field for checksum")
		}
		_, err = cheksumField.Write([]byte(digest))
		if err != nil {
			return nil, errors.Wrap(err, "Could not write checksum into request body")
		}
	}

	{
		fileField, err = mpWriter.CreateFormFile("file", "file.zip")
		if err != nil {
			return nil, errors.Wrap(err, "Could not create form field for file")
		}
		_, err = io.Copy(fileField, fileFieldRead)
		if err != nil {
			return nil, errors.Wrap(err, "Could not write the file content into form field")
		}
	}

	err = mpWriter.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Could not close the writer of request body")
	}

	req.contentTypeWithBoundary = mpWriter.FormDataContentType()

	return req, nil
}

func (req *Request) Send(method, url string) (*Response, error) {
	var (
		httpReq *http.Request
		err     error
	)
	httpReq, err = http.NewRequest(method, url, req.body)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create http request instance")
	}

	httpReq.Header.Set("Content-Type", req.contentTypeWithBoundary)
	httpReq.Header.Set("Content-Length", fmt.Sprintf("%d", req.body.Len()))
	// req.Header.Set("Authorization", "Bearer")

	httpResponse, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed on sending the request")
	}
	res := Response{}
	err = res.DeserializeResponse(httpResponse)
	if err != nil {
		return nil, errors.Wrap(err, "failed on parsing the response body")
	}

	return &res, nil
}
