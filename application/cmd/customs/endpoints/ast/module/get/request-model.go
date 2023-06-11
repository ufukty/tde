package ast

//go:generate serdeser request-model.go
type Request struct {
	ArchiveId string `json:"archive_id"`
}
