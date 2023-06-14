package module

//go:generate serdeser bind.go

type (
	Request struct {
		ArchiveId string `json:"archive_id"`
	}
)
