package endpoints_test

import (
	"tde/cmd/customs/endpoints"
	"tde/cmd/customs/endpoints/utilities"
	"tde/cmd/customs/endpoints/volmng"

	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func Test_UploadDownload(t *testing.T) {
	var (
		archiveId string
		checksum  string
		testFile  = "testdata/to-upload.zip"
	)

	tmp, err := os.MkdirTemp(os.TempDir(), "tde-customs-testing-endpoints-*")
	if err != nil {
		log.Fatalln(errors.Wrapf(err, "prep"))
	}
	var vm = volmng.NewVolumeManager(tmp)
	log.Printf("Using %s for volume manager mount", tmp)
	var handlers = endpoints.NewManager(vm)

	t.Run("Prep", func(t *testing.T) {
		fh, err := os.Open(testFile)
		if err != nil {
			log.Fatalln(errors.Wrapf(err, "prep, cheksum"))
		}
		defer fh.Close()
		if checksum, err = utilities.MD5(fh); err != nil {
			t.Error(errors.Wrapf(err, ""))
		}
	})

	t.Run("Uploading", func(t *testing.T) {
		var (
			err                error
			fileContent        []byte
			req                *http.Request
			expectedBodyRegexp *regexp.Regexp
			bindRes            endpoints.UploadResponse
			resrec             *httptest.ResponseRecorder
		)

		{
			fileContent, err = os.ReadFile(testFile)
			if err != nil {
				t.Fatal(errors.Wrapf(err, "prep"))
			}

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			md5Field, err := writer.CreateFormField("md5sum")
			if err != nil {
				t.Fatal(errors.Wrapf(err, "prep, create form field md5sum"))
			}
			md5Field.Write([]byte(checksum))
			fileField, err := writer.CreateFormFile("file", "file.zip")
			if err != nil {
				t.Fatal(errors.Wrapf(err, "prep, create form field file"))
			}
			fileField.Write(fileContent)
			writer.Close()
			contentType := writer.FormDataContentType()
			req, err = http.NewRequest("POST", "/module", body)
			if err != nil {
				t.Fatalf("prep. Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", contentType)
			req.Header.Set("Content-Length", "101010")
			req.Header.Set("Authorization", "Bearer")
			resrec = httptest.NewRecorder()
		}

		handlers.HandleUpload(resrec, req)
		var res = resrec.Result()

		{
			if res.StatusCode != http.StatusOK {
				t.Fatalf("validation/header. Got %v, want %v", resrec.Code, http.StatusOK)
			}
			var expectedResponseBody = `^{"archive_id":"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"}\n$`
			expectedBodyRegexp, err = regexp.Compile(expectedResponseBody)
			if err != nil {
				t.Fatal(errors.Wrapf(err, "validation prep"))
			}
			if !expectedBodyRegexp.Match(resrec.Body.Bytes()) {
				t.Fatalf("validation/body. Got %s, want %s", resrec.Body.Bytes(), expectedResponseBody)
			}
			bindRes = endpoints.UploadResponse{}
			if err = bindRes.DeserializeResponse(res); err != nil {
				t.Fatal(errors.Wrapf(err, "validation prep - deserialize"))
			}
			if bindRes.ArchiveID == "" {
				t.Fatal("validation. archive id is empty")
			}
			archiveId = bindRes.ArchiveID
		}
	})

	fmt.Println("Archive ID:", archiveId)

	t.Run("Downloading", func(t *testing.T) {
		var (
			r      = httptest.NewRequest("GET", "/module", nil)
			resrec = httptest.NewRecorder()
			w      *http.Response
			body   []byte
		)

		r = mux.SetURLVars(r, map[string]string{
			"id": archiveId,
		})

		r.Header.Set("Authorization", "Bearer")
		handlers.HandleDownload(resrec, r)
		w = resrec.Result()

		if w.StatusCode != http.StatusOK {
			t.Fatalf("validation. status code got %d", w.StatusCode)
		}

		var err error
		body, err = io.ReadAll(w.Body)
		if err != nil {
			t.Fatal("validation. expected a body")
		}
		if n := len(body); n == 0 {
			t.Fatal("validation. body size = 0")
		} else {
			fmt.Printf("Downloaded file size %d\n", n)
		}

		if got, err := utilities.MD5(bytes.NewReader(body)); err != nil {
			t.Fatal(errors.Wrap(err, "validation. request body checksum"))
		} else if got != checksum {
			t.Fatalf("validation checksum. got %s != expected %s", got, checksum)
		}

	})
}
