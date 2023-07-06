package context

import "tde/internal/cfg/context-resolution/context"

//go:generate serdeser bind.go
type Request struct {
	ArchiveId string `json:"archive_id"`
	Folder    string `json:"folder"`
	File      string `json:"file"`
	Function  string `json:"function"`
}

type Response struct {
	Context *context.Context `json:"context"`
}
