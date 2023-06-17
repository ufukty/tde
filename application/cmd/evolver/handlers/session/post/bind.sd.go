package session_post

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
func (req *Request) Send(method, url string) (*Response, error) {
	httpRequest, err := req.NewRequest(method, url)
	if err != nil {
		return nil, errors.Wrap(err, "failed on creating an object for request")
	}
	httpResponse, err := http.DefaultClient.Do(httpRequest)
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
func (res *Response) SerializeIntoResponseWriter(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		return errors.Wrap(err, "failed on serialization")
	}
	return nil
}
func (s *Response) DeserializeResponse(res *http.Response) error {
	err := json.NewDecoder(res.Body).Decode(s)
	if err != nil {
		return errors.Wrap(err, "failed on parsing the response body")
	}
	return nil
}
