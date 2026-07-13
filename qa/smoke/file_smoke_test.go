package smoke

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	emailapi "phishing-platform/internal/api/email"
)

func TestFileAnalysisSmoke(t *testing.T) {

	setupDatabase()

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile(
		"file",
		"sample.txt",
	)

	if err != nil {

		t.Fatal(err)

	}

	_, err = part.Write([]byte(
		"This is a smoke test file.",
	))

	if err != nil {

		t.Fatal(err)

	}

	writer.Close()

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/analyze-file",
		&body,
	)

	req.Header.Set(
		"Content-Type",
		writer.FormDataContentType(),
	)

	rec := httptest.NewRecorder()

	emailapi.AnalyzeFileHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {

		t.Fatalf(
			"Expected 200 OK, got %d",
			rec.Code,
		)

	}

	t.Log("==============================")
	t.Log("FILE ANALYSIS SMOKE TEST PASSED")
	t.Log("File Upload Working")
	t.Log("Analysis Engine Working")
	t.Log("==============================")
}
