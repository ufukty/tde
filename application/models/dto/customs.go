package dto

//go:generate serdeser customs.go

type (
	Customs_Upload_Request struct {
		Token string `json:"token"`
	}

	Customs_Upload_Response struct {
		ArchiveID string `json:"archive_id"`
	}
)
