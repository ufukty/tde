package list_packages

//go:generate serdeser bind.go

type Request struct {
	ArchiveId string `json:"archive_id"`
}

type Package struct {
	Folder     string `json:"folder"`
	ImportPath string `json:"import_path"`
}

type Response struct {
	PackageList []Package `json:"package"`
}
