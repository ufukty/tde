package endpoints

import (
	"tde/cmd/customs/endpoints/utilities"
	"tde/config"
	"tde/internal/folders/archive"

	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

//go:generate serdeser module-upload.go

// NOTE: DON'T use serdeser for multipart/form-data requests
type UploadRequest struct {
	body                    *bytes.Buffer
	contentTypeWithBoundary string
}

type UploadResponse struct {
	ArchiveID string `json:"archive_id"`
}

func NewUploadRequest(file io.Reader) (*UploadRequest, error) {
	var (
		err          error
		digest       string
		fileField    io.Writer
		cheksumField io.Writer
		req          = new(UploadRequest)
		mpWriter     *multipart.Writer
	)

	req.body = bytes.NewBuffer([]byte{})
	mpWriter = multipart.NewWriter(req.body)

	var fileFieldRead = bytes.NewBuffer([]byte{})
	var checksumFieldRead = io.TeeReader(file, fileFieldRead)

	{
		digest, err = utilities.MD5(checksumFieldRead)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to find checksum of file")
		}
		cheksumField, err = mpWriter.CreateFormField("md5sum")
		if err != nil {
			return nil, errors.Wrap(err, "Could not create form field for checksum")
		}
		_, err = cheksumField.Write([]byte(digest))
		if err != nil {
			return nil, errors.Wrap(err, "Could not write checksum into request body")
		}
	}

	{
		fileField, err = mpWriter.CreateFormFile("file", "file.zip")
		if err != nil {
			return nil, errors.Wrap(err, "Could not create form field for file")
		}
		_, err = io.Copy(fileField, fileFieldRead)
		if err != nil {
			return nil, errors.Wrap(err, "Could not write the file content into form field")
		}
	}

	err = mpWriter.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Could not close the writer of request body")
	}

	req.contentTypeWithBoundary = mpWriter.FormDataContentType()

	return req, nil
}

func (req *UploadRequest) Send() (*UploadResponse, error) {
	var (
		httpReq *http.Request
		err     error
	)
	httpReq, err = http.NewRequest(config.CustomsModuleUpload.Method.String(), config.CustomsModuleUpload.Url(), req.body)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create http request instance")
	}

	httpReq.Header.Set("Content-Type", req.contentTypeWithBoundary)
	httpReq.Header.Set("Content-Length", fmt.Sprintf("%d", req.body.Len()))
	// req.Header.Set("Authorization", "Bearer")

	httpResponse, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed on sending the request")
	}
	if httpResponse.StatusCode != http.StatusOK {
		buf := bytes.NewBuffer([]byte{})
		_, err := io.Copy(buf, httpResponse.Body)
		if err != nil {
			return nil, errors.Wrap(err, "Could not read response body")
		}
		return nil, errors.New(buf.String())
	}

	res := UploadResponse{}
	err = res.DeserializeResponse(httpResponse)
	if err != nil {
		return nil, errors.Wrap(err, "failed on parsing the response body")
	}

	return &res, nil
}

func (em EndpointsManager) UploadHandler() func(w http.ResponseWriter, r *http.Request) {
	const (
		MAX_CONTENT_LENGTH = 40 * 1024 * 1024
		ALLOWED_MIME_TYPE  = "multipart/form-data"
	)

	var processUploadWithVolumeManager = func(em *EndpointsManager, r *http.Request) (archiveId string, errResp string, err error) {
		fh, _, err := r.FormFile("file")
		if err != nil {
			return "", "", errors.Wrap(err, "retrieving form part 'file'")
		}
		defer fh.Close()

		archiveId, errFile, err := em.vm.New(fh)
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

	var checkHeaderContentType = func(r *http.Request) error {
		if contentTypeHeader := r.Header.Get("Content-Type")[:len(ALLOWED_MIME_TYPE)]; contentTypeHeader != ALLOWED_MIME_TYPE {
			return errors.New(fmt.Sprintf("Content-Type '%s' is not allowed.", contentTypeHeader))
		}
		return nil
	}

	var checkHeaderContentLength = func(r *http.Request) error {
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

	var checkMD5Sum = func(r *http.Request) error {
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

	return func(w http.ResponseWriter, r *http.Request) {
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

		archiveID, responseText, err = processUploadWithVolumeManager(&em, r)
		if err != nil {
			var message = "Could not process the upload response text: " + responseText
			log.Println(errors.Wrap(errors.Wrap(err, message), requestID))
			http.Error(w, responseText, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", mime.TypeByExtension("json"))
		w.WriteHeader(http.StatusOK)

		var res = &UploadResponse{
			ArchiveID: archiveID,
		}

		res.SerializeIntoResponseWriter(w)
	}
}
