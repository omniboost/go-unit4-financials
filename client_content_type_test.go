package financials_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	financials "github.com/omniboost/go-unit4-financials"
)

func fakeResponse(status int, contentType, body string) *http.Response {
	return &http.Response{
		StatusCode:    status,
		Status:        http.StatusText(status),
		Header:        http.Header{"Content-Type": []string{contentType}},
		Body:          io.NopCloser(strings.NewReader(body)),
		ContentLength: int64(len(body)),
	}
}

func TestCheckResponseUnauthorizedNoBody(t *testing.T) {
	resp := fakeResponse(401, "", "")

	err := financials.CheckResponse(resp)
	if err == nil {
		t.Fatal("expected an error")
	}

	if err.Error() != "Unauthorized" {
		t.Fatalf("expected the status as error message, got: %s", err.Error())
	}
}

func TestCheckResponseTextHTML(t *testing.T) {
	resp := fakeResponse(502, "text/html; charset=utf-8", "<html><body>Bad Gateway</body></html>")

	err := financials.CheckResponse(resp)
	if err == nil {
		t.Fatal("expected an error")
	}

	if strings.Contains(err.Error(), "<html>") {
		t.Fatalf("expected the html body not to leak into the error message, got: %s", err.Error())
	}
}

func TestCheckResponseTextPlain(t *testing.T) {
	resp := fakeResponse(500, "text/plain; charset=utf-8", "Internal Server Error: something went wrong")

	err := financials.CheckResponse(resp)
	if err == nil {
		t.Fatal("expected an error")
	}

	if !strings.Contains(err.Error(), "something went wrong") {
		t.Fatalf("expected the plain text body to be part of the error message, got: %s", err.Error())
	}
}
