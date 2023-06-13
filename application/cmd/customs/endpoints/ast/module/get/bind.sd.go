package ast

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func (req *Request) NewRequest(method, url string) (*http.Request, error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed on serialization")
	}
	httpRequest, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, errors.Wrap(err, "failed on creating request object")
	}
	return httpRequest, nil
}
func (req *Request) ParseRequest(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return errors.Wrap(err, "failed on parsing the response body")
	}
	return nil
}
