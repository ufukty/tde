package validator

import (
	"tde/internal/microservices/errors/bucket"
	"tde/internal/microservices/errors/detailed"
	"tde/internal/microservices/validator/rule"

	"net/http"
)

func CheckHeaders(r *http.Request, headerRules *map[string]*rule.Rule) *bucket.Bucket {
	var mainBucket = bucket.New()
	for headerKey, rule := range *headerRules {
		if bucket := rule.Test(r.Form.Get(headerKey)); bucket != nil {
			mainBucket.Add(bucket.Tag(detailed.New("Invalid value for header field "+headerKey, "@CheckHeaders")))
		}
	}
	return mainBucket
}
