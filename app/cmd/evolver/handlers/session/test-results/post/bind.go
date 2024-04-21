package test_results_post

//go:generate serdeser bind.go

type TestResult struct {
	Sid       string
	Completed bool // without run time error
	Distance  float64
}

type Request struct {
	TestResults []TestResult
	CaseID      string
}

type EvolverService_Results_Response struct {
	Success bool
}
