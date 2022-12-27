package in_program_models

type DiscoveryResponse struct {
	ModuleAbsolutePath      string
	TargetPackageImportPath string
}

type TestFunctionDetails struct {
	TestFunctionName         string
	TestFunctionFilename     string
	TestFunctionStartingLine int
	TestFunctionEndingLine   int
}

type Config struct {
	TestFunctionDetails *TestFunctionDetails
	DiscoveredDetails   *DiscoveryResponse
}
