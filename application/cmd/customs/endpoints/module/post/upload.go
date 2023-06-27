package module

import (
	"tde/cmd/customs/internal/utilities"
	volume_manager "tde/cmd/customs/internal/volume-manager"
	"tde/internal/folders/archive"
	"tde/internal/microservices/logger"

	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	MAX_CONTENT_LENGTH = 40 * 1024 * 1024
	ALLOWED_MIME_TYPE  = "multipart/form-data"
)

var (
	log           = logger.NewLogger("http handler in customs/module/post")
	volumeManager *volume_manager.VolumeManager
)

func RegisterVolumeManager(vm *volume_manager.VolumeManager) {
	volumeManager = vm
}

func writeZipIntoDest(r *http.Request, destForOriginal string) error {
	srcFileHandler, _, err := r.FormFile("file")
	if err != nil {
		return errors.Wrap(err, "Could not get the file from request")
	}
	defer srcFileHandler.Close()

	dest, err := os.Create(destForOriginal)
	if err != nil {
		return errors.Wrap(err, "Create destination file")
	}
	defer dest.Close()

	_, err = io.Copy(dest, srcFileHandler)
	if err != nil {
		return errors.Wrap(err, "Write into destination file")
	}
	return nil
}

func checkHeaderContentType(r *http.Request) error {
	if contentTypeHeader := r.Header.Get("Content-Type")[:len(ALLOWED_MIME_TYPE)]; contentTypeHeader != ALLOWED_MIME_TYPE {
		return errors.New(fmt.Sprintf("Content-Type '%s' is not allowed.", contentTypeHeader))
	}
	return nil
}

func checkHeaderContentLength(r *http.Request) error {
	var (
		contentLength    int
		contentLengthStr string
		err              error
	)
	contentLengthStr = r.Header.Get("Content-Length")
	if contentLengthStr == "" {
		return errors.New("Content-Length is empty")
	}
	contentLength, err = strconv.Atoi(contentLengthStr)
	if err != nil {
		return errors.New("Content-Length is not an integer")
	}
	if contentLength > MAX_CONTENT_LENGTH {
		return errors.New("Content-Length is bigger than allowed")
	}
	return nil
}

func checkMD5Sum(r *http.Request) error {
	var (
		err              error
		filePart         multipart.File
		md5sumSent       string
		md5sumCalculated string
	)

	md5sumSent = r.FormValue("md5sum")
	if md5sumSent == "" {
		return errors.New("Error MD5 checksum not found")
	}

	filePart, _, err = r.FormFile("file")
	if err != nil {
		return errors.Wrap(err, "Error retrieving file")
	}
	defer filePart.Close()
	md5sumCalculated, err = utilities.MD5(filePart)
	if err != nil {
		return errors.Wrap(err, "could not call the md5 utility")
	}
	if md5sumCalculated != md5sumSent {
		return errors.New("MD5 checksum mismatch")
	}
	return nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		archiveID       string
		err             error
		requestID       string
		destForOriginal string
		destForExtract  string
	)

	requestID = uuid.NewString()

	if err = checkHeaderContentType(r); err != nil {
		var message = "Unaccepted Content-Type header"
		log.Println(errors.Wrap(err, message))
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	if err = checkHeaderContentLength(r); err != nil {
		var message = "Unaccepted Content-Length header"
		log.Println(errors.Wrap(err, message))
		http.Error(w, message, http.StatusRequestEntityTooLarge)
		return
	}

	if err = r.ParseMultipartForm(0); err != nil {
		var message = "Error parsing form data"
		log.Println(errors.Wrap(errors.Wrap(err, message), requestID))
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	archiveID = volumeManager.CreateUniqueFilename()
	destForOriginal, destForExtract, err = volumeManager.CreateDestPath(archiveID)
	if err != nil {
		var message = "Could not create dir entries to place incoming file into"
		log.Println(errors.Wrap(errors.Wrap(err, message), requestID))
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	if err = checkMD5Sum(r); err != nil {
		var message = "Could not validate check sum of file"
		log.Println(errors.Wrap(errors.Wrap(err, message), requestID))
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	err = writeZipIntoDest(r, destForOriginal)
	if err != nil {
		var message = "Could not parse the file part of multipart request"
		log.Println(errors.Wrap(errors.Wrap(err, message), requestID))
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	if err = archive.Unarchive(destForOriginal, destForExtract); err != nil {
		var message = "Could not process the upload"
		log.Println(errors.Wrap(errors.Wrap(err, message), requestID))
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	var res = &Response{
		ArchiveID: archiveID,
	}

	res.SerializeIntoResponseWriter(w)
}
