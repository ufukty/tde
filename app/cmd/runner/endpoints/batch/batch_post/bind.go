package batch_post

import "go/ast"

//go:generate serdeser bind.go

type (
	Subject struct {
		Sid      string
		FuncDecl *ast.FuncDecl `json:"func_decl"`
	}

	Request struct {
		Subjects     []Subject
		ArchiveID    string
		FileTemplate *ast.File `json:"file"`
		CaseID       string
	}

	Response struct {
		Registered bool
	}
)
