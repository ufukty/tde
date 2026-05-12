package i18n

import "fmt"

var (
	ErrAstConversionFailed   = fmt.Errorf("AST convertion has failed")
	ErrMissingParameter      = fmt.Errorf("One or more required parameters are not found in the request; either in url, header or body.")
	ErrMultiplePackagesFound = fmt.Errorf("More than 1 package found at the directory")
	ErrNoPackagesFound       = fmt.Errorf("Package not found in directory")
	ErrInputSanitization     = fmt.Errorf("One or more inputs are invalid. Try again after fixing them.")
	ErrFileNotFoundInPackage = fmt.Errorf("Specified package doesn't contain the given file.")
)
