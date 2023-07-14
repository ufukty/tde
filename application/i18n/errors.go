package i18n

import "errors"

var (
	ErrAstConversionFailed   = errors.New("AST convertion has failed")
	ErrMissingParameter      = errors.New("One or more required parameters are not found in the request; either in url, header or body.")
	ErrMultiplePackagesFound = errors.New("More than 1 package found at the directory")
	ErrNoPackagesFound       = errors.New("Package not found in directory")
	ErrInputSanitization     = errors.New("One or more inputs are invalid. Try again after fixing them.")
	ErrFileNotFoundInPackage = errors.New("Specified package doesn't contain the given file.")
)
