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

	RunnerService_Batch_Request struct {
		Candidates   []Candidate
		ArchiveID    string
		FileTemplate *ast.File `json:"file"`
		CaseID       string
	}

	RunnerService_Batch_Response struct {
		Registered bool
	}
)

type (
	TestResult struct {
		CandidateID string
		Completed   bool // without run time error
		Distance    float64
	}

	EvolverService_Results_Request struct {
		TestResults []TestResult
		CaseID      string
	}

	EvolverService_Results_Response struct {
		Success bool
	}
)
