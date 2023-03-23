package models

import (
	"fmt"
	"tde/models/in_program_models"
)

type EmbeddingConfig struct {
	ImplementationPackagePath string // eg. .../userPackage
	ImplementationFilePath    string // eg. .../userPackage/userFile.go
	TestFunctionPath          string // eg. .../userPackage/userFile_tde.go
	TestFunctionName          string // eg. TDE_WordReverse
	PackageImportPath         string // eg. userModule/path/to/userPackageDir/userPackage
	TesterHeadDirPath         string // eg. .../userPackage/tde
	TesterHeadFilePath        string // eg. .../userPackage/tde/main_tde.go
}

func NewEmbeddingConfig(implementationPackageFile, implementationFilePath, testFunctionPath, targetPackageImportPath, testFunctionName string) *EmbeddingConfig {
	return &EmbeddingConfig{
		ImplementationPackagePath: implementationPackageFile,
		ImplementationFilePath:    implementationFilePath,
		TestFunctionPath:          testFunctionPath,
		TestFunctionName:          testFunctionName,
		PackageImportPath:         targetPackageImportPath,
		TesterHeadDirPath:         fmt.Sprintf("%s/tde", implementationPackageFile),
		TesterHeadFilePath:        fmt.Sprintf("%s/tde/main_tde.go", implementationPackageFile),
	}
}

type OperationConfig struct {
	ModulePath    string
	TempDir       string
	CandidateDirs map[in_program_models.CandidateID]string
}
