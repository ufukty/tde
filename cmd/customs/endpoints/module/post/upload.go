package module

import (
	"log"
	"strconv"
	"tde/cmd/customs/internal/volume-manager"
	"tde/models/dto"

	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// type Error struct {
// 	public  []error
// 	private []error
// }

// func NewError() *Error {
// 	return &Error{}
// }

// func (e *Error) Wrap(pub, pri error) {
// 	if pub != nil {
// 		e.public = append(e.public, pub)
// 	}
// 	if pri != nil {
// 		e.private = append(e.private, pri)
// 	}
// }

// func (e *Error) IsIn(pub, pri error) bool {
// 	if pub != nil {
// 		return slices.Index(e.public, pub) != -1
// 	}
// 	if pri != nil {
// 		return slices.Index(e.private, pri) != -1
// 	}
// 	return false
// }

const (
	HTTPErrBadRequest = "Bad request"
)

const (
	maxAllowedContentLength = 40 * 1024 * 1024
)

var volumeManager *volume_manager.VolumeManager

func RegisterVolumeManager(vm *volume_manager.VolumeManager) {
	volumeManager = vm
}

func jsonPart(part *multipart.Part, req *dto.Customs_Upload_Request) error {
	err := json.NewDecoder(part).Decode(req)
	if err != nil {
		return errors.Wrap(err, "Deserialization")
	}
	return nil
}

func zipPart(part *multipart.Part, storagePath string) error {
	dest, err := os.Create(storagePath)
	if err != nil {
		return errors.Wrap(err, "Create destination file")
	}
	defer dest.Close()

	_, err = io.Copy(dest, part)
	if err != nil {
		return errors.Wrap(err, "Write into destination file")
	}
	return nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		archiveID        string
		contentLength    int
		contentLengthStr string
		err              error
		part             *multipart.Part
		req              *dto.Customs_Upload_Request
		requestID        string
		storagePath      string
	)

	requestID = uuid.NewString()

	contentLengthStr = r.Header.Get("Content-Length")
	if contentLengthStr == "" {
		log.Println(requestID, `Content-Length is empty`, err)
		http.Error(w, HTTPErrBadRequest, http.StatusRequestEntityTooLarge)
		return
	}
	contentLength, err = strconv.Atoi(contentLengthStr)
	if err != nil {
		log.Println(requestID, `strconv.Atoi(r.Header.Get("Content-Length"))`, err)
		http.Error(w, HTTPErrBadRequest, http.StatusRequestEntityTooLarge)
		return
	}
	if contentLength > maxAllowedContentLength {
		log.Println(requestID, `contentLength > maxAllowedContentLength`, err)
		http.Error(w, HTTPErrBadRequest, http.StatusRequestEntityTooLarge)
		return
	}

	req = &dto.Customs_Upload_Request{}

	archiveID = volumeManager.CreateUniqueFilename()
	storagePath, err = volumeManager.CreateDestPath(archiveID)
	if err != nil {
		http.Error(w, errors.Wrap(err, "Could not create dir entries to place incoming file into").Error(), http.StatusInternalServerError)
		return
	}

	multipartReader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, errors.Wrap(err, "").Error(), http.StatusBadRequest)
		return
	}
	for {
		part, err = multipartReader.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			}
			http.Error(w, errors.Wrap(err, "Could not complete parsing multipart request").Error(), http.StatusBadRequest)
			return
		}
		switch part.FormName() {
		case "json":
			err = jsonPart(part, req)
			if err != nil {
				http.Error(w, errors.Wrap(err, "Could not parse the file part of multipart request").Error(), http.StatusBadRequest)
				return
			}
		case "file":
			err = zipPart(part, filepath.Join(storagePath, archiveID+".zip"))
			if err != nil {
				http.Error(w, errors.Wrap(err, "Could not parse the file part of multipart request").Error(), http.StatusBadRequest)
				return
			}
		}
	}

	var res = &dto.Customs_Upload_Response{
		ArchiveID: archiveID,
	}

	res.SerializeIntoResponseWriter(w)
}
