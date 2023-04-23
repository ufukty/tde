package models

import (
	"go/ast"
)

//go:generate serdeser runner_service_new_test_endpoint.go

type Candidate struct {
	CandidateID string
	FuncDecl    *ast.FuncDecl `json:"func_decl"`
}

type TestResult struct {
	Completed bool
	Distance  float64
}

type RunnerService_NewTest_Request struct {
	Candidates   []Candidate
	ArchiveID    string
	FileTemplate *ast.File `json:"file"`
}

type RunnerService_NewTest_Response struct {
	CandidateId string
	TestResults []TestResult
}

// type RunnerService interface {
// 	NewTest(RunnerService_NewTest_Request) RunnerService_NewTest_Response // POST /test
// }
