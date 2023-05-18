package module

import (
	"tde/models/dto"

	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

func Test_Endpoint(t *testing.T) {

	req := dto.Customs_Upload_Request{
		Token: "922a5105-28e4-55bb-bbf5-1d3bb72ec38d",
	}
	httpReq, err := req.NewRequest("POST", "http://localhost:8080")
	if err != nil {
		t.Error(errors.Wrapf(err, ""))
	}

	// Set up the zip file
	file, err := os.Open("file.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Set up the HTTP request
	requestBodyBuf := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(requestBodyBuf)
	zipWriter, err := multipartWriter.CreateFormFile("file", "file.zip")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(zipWriter, file)
	if err != nil {
		panic(err)
	}
	multipartWriter.Close()

	httpReq.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// Send the HTTP request
	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		panic(err)
	}
	defer httpResp.Body.Close()

	res := dto.Customs_Upload_Response{}
	res.DeserializeResponse(httpResp)
	spew.Println(res)
}
