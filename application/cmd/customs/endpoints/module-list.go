package endpoints

import "net/http"

//go:generate serdeser module-list.go

type ListRequest struct {
	ArchiveId string `json:"archive_id"`
}

type ListRequestPackage struct {
	Folder     string `json:"folder"`
	ImportPath string `json:"import_path"`
}

type HandleListResponse struct {
	PackageList []ListRequestPackage `json:"package"`
}

func (em EndpointsManager) ListHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}

}
