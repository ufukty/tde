package module

//go:generate serdeser dto.go

type (
	Request struct {
		ArchiveId string `json:"archive_id"`
	}
)
