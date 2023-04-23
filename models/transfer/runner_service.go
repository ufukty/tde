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

	TestResult struct {
		Completed bool
		Distance  float64
	}

	RunnerService_NewTest_Request struct {
		Candidates   []Candidate
		ArchiveID    string
		FileTemplate *ast.File `json:"file"`
	}

	RunnerService_NewTest_Response struct {
		CandidateId string
		TestResults []TestResult
	}
)
