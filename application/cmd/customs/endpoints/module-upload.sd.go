package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func (res *UploadResponse) SerializeIntoResponseWriter(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		return errors.Wrap(err, "failed on serialization")
	}
	return nil
}
func (s *UploadResponse) DeserializeResponse(res *http.Response) error {
	err := json.NewDecoder(res.Body).Decode(s)
	if err != nil {
		return errors.Wrap(err, "failed on parsing the response body")
	}
	return nil
}
