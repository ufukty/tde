package module

//go:generate serdeser bind.go

type Request struct {
	Token string `json:"token"`
}

type Response struct {
	ArchiveID string `json:"archive_id"`
}
