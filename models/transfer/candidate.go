package models

type Candidate struct {
	CandidateID        string
	FuncDeclSerialized string
}

type TestResult struct {
	Completed bool
	Distance  float64
}

type RunnerServiceNewTestRequest struct {
	Candidates             []Candidate
	ArchiveID              string
	FileTemplateSerialized string
}

type RunnerServiceNewTestResponse struct {
	CandidateId string
	TestResults []TestResult
}
