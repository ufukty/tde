package models

import (
	"bytes"
	"encoding/json"
	"net/http"
	"github.com/pkg/errors"
)

func (s *RunnerService_NewTest_Request) NewRequest(method, url string) (*http.Request, error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(s)
	if err != nil {
		return nil, errors.Wrap(err, "failed on serialization")
	}
	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, errors.Wrap(err, "failed on creating request object")
	}
	return req, nil
}
func (s *RunnerService_NewTest_Request) ParseRequest(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		return errors.Wrap(err, "failed on serialization")
	}
	return nil
}
func (s *RunnerService_NewTest_Response) SerializeIntoResponseWriter(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		return errors.Wrap(err, "failed on serialization")
	}
	return nil
}
func (s *RunnerService_NewTest_Response) DeserializeResponse(res *http.Response) error {
	err := json.NewDecoder(res.Body).Decode(s)
	if err != nil {
		return errors.Wrap(err, "failed on serialization")
	}
	return nil
}
