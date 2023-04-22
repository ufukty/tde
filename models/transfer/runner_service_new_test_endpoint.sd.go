package models

import (
	"encoding/json"
	"net/http"
	"github.com/pkg/errors"
)

func (s *RunnerService_NewTest_Request) Serialize(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		return errors.Wrap(err, "failed on serialization")
	}
	return nil
}
func (s *RunnerService_NewTest_Request) Deserialize(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		return errors.Wrap(err, "failed on serialization")
	}
	return nil
}
func (s *RunnerService_NewTest_Response) Serialize(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		return errors.Wrap(err, "failed on serialization")
	}
	return nil
}
func (s *RunnerService_NewTest_Response) Deserialize(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		return errors.Wrap(err, "failed on serialization")
	}
	return nil
}
