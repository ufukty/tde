package session_post

//go:generate serdeser bind.go

type Request struct {
	ArchiveId string `json:"archive_id"`
	// SnapshotId string `json:"snapshot_id"` // TODO:
	Iterations int `json:"iterations"`
	Population int `json:"population"`

	Package int `json:"package"`
}

type Response struct{}
