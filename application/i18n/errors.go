package i18n

import "errors"

// To be used by Customs service
var (
	ErrAstConversionFailed   = errors.New("AST convertion has failed")
	ErrMissingParameter      = errors.New("One or required parameters are not found in the request; either in url, header or body.")
	ErrMultiplePackagesFound = errors.New("More than 1 package found at the directory")
	ErrNoPackagesFound       = errors.New("Package not found in directory")
)
