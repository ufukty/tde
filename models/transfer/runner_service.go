package models

import (
	"go/ast"
)

//go:generate serdeser runner_service.go

type (
	Candidate struct {
		CandidateID string
		FuncDecl    *ast.FuncDecl `json:"func_decl"`
	}

	RunnerService_NewTest_Request struct {
		Candidates   []Candidate
		ArchiveID    string
		FileTemplate *ast.File `json:"file"`
		CaseID       string
	}

	RunnerService_NewTest_Response struct {
		Registered bool
	}
)

type (
	TestResult struct {
		CandidateID string
		Completed   bool // without run time error
		Distance    float64
	}

	EvolverService_TestResultsAcceptance_Request struct {
		TestResults []TestResult
		CaseID      string
	}

	EvolverService_TestResultsAcceptance_Response struct {
		Success bool
	}
)
