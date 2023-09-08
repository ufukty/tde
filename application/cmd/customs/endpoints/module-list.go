package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"tde/internal/evolution/evaluation/list"
	"tde/internal/microservices/utilities"

	"github.com/google/uuid"
)

type ListRequest struct {
	ArchiveId string `url:"id"`
	// Folder    string `body:"folder"`
	// ImportPath string `body:"import_path"`
}

type ListResponse struct {
	Results list.Packages `json:"package"`
}

func (em EndpointsManager) ListHandler() func(w http.ResponseWriter, r *http.Request) {
	var sanitizeRequest = func(bq *ListRequest) error {
		if typedArchiveId, err := uuid.Parse(bq.ArchiveId); err != nil {
			return fmt.Errorf("invalid uuid for archiveid %q: %w", bq.ArchiveId, err)
		} else {
			bq.ArchiveId = typedArchiveId.String()
		}

		return nil
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			pkgs list.Packages
			bq   *ListRequest
			err  error
		)

		if bq, err = utilities.ParseRequest[ListRequest](r); err != nil {
			log.Println(fmt.Errorf("bad request: %w", err))
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if err = sanitizeRequest(bq); err != nil {
			log.Println(fmt.Errorf("sanitize: %w", err))
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if bundle, zip, extract := em.vm.CheckIfExists(bq.ArchiveId); !(bundle && zip && extract) {
			log.Println(fmt.Errorf("requested archive doesn't exist %q", bq.ArchiveId))
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		var _, _, extract = em.vm.FindPath(bq.ArchiveId)
		if pkgs, err = list.ListAllPackages(extract); err != nil {
			log.Println(fmt.Errorf("listing packages in dir %q: %w", extract, err))
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		var bs = &ListResponse{
			Results: pkgs,
		}
		if err = utilities.WriteJsonResponse(bs, w); err != nil {
			log.Println(fmt.Errorf("writing response body: %w", err))
		}
	}
}
