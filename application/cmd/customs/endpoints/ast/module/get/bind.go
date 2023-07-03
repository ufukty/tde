package ast

import "go/ast"

//go:generate serdeser bind.go
type Request struct {
	ArchiveId string `json:"archive_id"`
	Folder    string `json:"folder"`
}

type Response struct {
	Package *ast.Package `json:"package"`
}
