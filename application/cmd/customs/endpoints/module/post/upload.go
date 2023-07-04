package post

import (
	"mime"
	"tde/cmd/customs/internal/utilities"
	volume_manager "tde/cmd/customs/internal/volume-manager"
	"tde/internal/folders/archive"
	"tde/internal/microservices/logger"

	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	MAX_CONTENT_LENGTH = 40 * 1024 * 1024
	ALLOWED_MIME_TYPE  = "multipart/form-data"
)

var (
	log = logger.NewLogger("http handler in customs/module/post")
	vm  *volume_manager.VolumeManager
)

func RegisterVolumeManager(vm_ *volume_manager.VolumeManager) {
	vm = vm_
}

func processUploadWithVolumeManager(r *http.Request) (archiveId string, errResp string, err error) {
	fh, _, err := r.FormFile("file")
	if err != nil {
		return "", "", errors.Wrap(err, "retrieving form part 'file'")
	}
	defer fh.Close()

	archiveId, errFile, err := vm.New(fh)
	if err != nil {
		if errors.Is(err, archive.ErrExtensionUnallowed) {
			errResp = "Check file extension"
		} else if errors.Is(err, archive.ErrZipFileExceedsLimit) {
			errResp = fmt.Sprintf("File exceeds size limit: %s", errFile)
		} else if errors.Is(err, archive.ErrZipExceedsLimit) {
			errResp = "Zip exceeds size limit"
		} else if errors.Is(err, archive.ErrRelativePathFound) {
			errResp = fmt.Sprintf("Parent refs in paths are unallowed: %s", errFile)
		} else if errors.Is(err, archive.ErrTooManyFiles) {
			errResp = "Zip contains too many files"
		} else if errors.Is(err, archive.ErrSubfolderExceedingDepth) {
			errResp = fmt.Sprintf("Path exceeds the depth limit '%s'", errFile)
		} else {
			errResp = "Woops"
		}
		return "", errResp, errors.Wrap(err, "passing to VolumeManager")
	}
	return archiveId, "", nil
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
		return errors.New("retrieving form part 'md5sum'")
	}

	filePart, _, err = r.FormFile("file")
	if err != nil {
		return errors.Wrap(err, "retrieving form part 'file'")
	}
	defer filePart.Close()
	md5sumCalculated, err = utilities.MD5(filePart)
	if err != nil {
		return errors.Wrap(err, "getting md5sum of zip")
	}
	if md5sumCalculated != md5sumSent {
		return errors.New("mismatched checksum")
	}
	return nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		requestID    string
		archiveID    string
		err          error
		responseText string
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

	if err = checkMD5Sum(r); err != nil {
		var message = "Could not validate check sum of file"
		log.Println(errors.Wrap(errors.Wrap(err, message), requestID))
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	archiveID, responseText, err = processUploadWithVolumeManager(r)
	if err != nil {
		var message = "Could not process the upload response text: " + responseText
		log.Println(errors.Wrap(errors.Wrap(err, message), requestID))
		http.Error(w, responseText, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", mime.TypeByExtension("json"))
	w.WriteHeader(http.StatusOK)

	var res = &Response{
		ArchiveID: archiveID,
	}

	res.SerializeIntoResponseWriter(w)
}
