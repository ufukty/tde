package batch_post

import "go/ast"

//go:generate serdeser bind.go

type (
	Candidate struct {
		CandidateID string
		FuncDecl    *ast.FuncDecl `json:"func_decl"`
	}

	Request struct {
		Candidates   []Candidate
		ArchiveID    string
		FileTemplate *ast.File `json:"file"`
		CaseID       string
	}

	Response struct {
		Registered bool
	}
)
