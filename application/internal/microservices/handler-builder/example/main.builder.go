package endpoint

import (
	"io"
	"net/http"
	"os"
)

func ParseHeader() {}

func Handler(w http.ResponseWriter, r *http.Request) {
	var req = Request{
		Header: RequestHeader{},
		Body:   RequestBody{},
	}
	r.ParseMultipartForm(4 * 1024 * 1024)
	req.Header.Authorization = r.FormValue("Authorization")
	var res = Endpoint(&req)
	w.Write(res.Body)
	w.WriteHeader(res.Header.StatusCode)
}

func Endpoint(r *Request) *Response {
	var fh, err = os.Create("dest")
	if err != nil {
		return nil
	}
	_, err = io.Copy(fh, r.Body.File)
	if err != nil {
		return &Response{
			Header: ResponseHeader{
				StatusCode: http.StatusInternalServerError,
			},
		}
	}

	return &Response{
		Header: ResponseHeader{
			StatusCode: http.StatusOK,
		},
	}
}
