package module

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	volume_manager "tde/cmd/customs/internal/volume-manager"
	"testing"

	"github.com/pkg/errors"
)

func init() {
	var root, err = os.MkdirTemp(os.TempDir(), "*")
	if err != nil {
		panic(errors.Wrapf(err, "prep"))
	}
	volumeManager = volume_manager.NewVolumeManager(root)
}

func TestUploadHandler(t *testing.T) {
	var (
		err                error
		fileContent        []byte
		req                *http.Request
		expectedBodyRegexp *regexp.Regexp
	)
	fileContent, err = os.ReadFile("test-files/do-not-edit")
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}
	var md5sum = "4b5f52abb39268d758e369fa535f5e80"
	var body = &bytes.Buffer{}
	var writer = multipart.NewWriter(body)
	var md5Field, _ = writer.CreateFormField("md5sum")
	md5Field.Write([]byte(md5sum))
	var fileField, _ = writer.CreateFormFile("file", "file.zip")
	fileField.Write(fileContent)
	writer.Close()
	var contentType = writer.FormDataContentType()
	req, err = http.NewRequest("POST", "/api/v1.0.0/customs/module", body)
	if err != nil {
		t.Fatalf("prep. Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", "101010")
	req.Header.Set("Authorization", "Bearer")
	var rr = httptest.NewRecorder()
	Handler(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("validation/header. Got %v, want %v", rr.Code, http.StatusOK)
	}
	var expectedResponseBody = `^{"archive_id":"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"}\n$`
	expectedBodyRegexp, err = regexp.Compile(expectedResponseBody)
	if err != nil {
		t.Error(errors.Wrapf(err, "validation prep"))
	}
	fmt.Printf("%v", rr.Body.Bytes())
	if !expectedBodyRegexp.Match(rr.Body.Bytes()) {
		t.Errorf("validation/body. Got %s, want %s", rr.Body.Bytes(), expectedResponseBody)
	}
}
