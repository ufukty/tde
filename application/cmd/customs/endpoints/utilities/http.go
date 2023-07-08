package utilities

import (
	"tde/internal/microservices/errors/bucket"
	"tde/internal/microservices/errors/detailed"

	"fmt"
	"net/http"
)

func AssertHeader(r *http.Request, headerField, want string) *detailed.DetailedError {
	var got = r.Header.Get(headerField)
	var expected = want
	if expected != got {
		return detailed.New(
			fmt.Sprintf("Value for %s header field is unexpected.", headerField),
			fmt.Sprintf("Expected %s, Got %s", expected, got),
		)
	}
	return nil
}

func CheckHeaders(r *http.Request) *bucket.Bucket {
	var errs = new(bucket.Bucket)
	if de := AssertHeader(r, "Content-Type", "application/json"); de != nil {
		errs.Add(de)
	}
	return errs
}
