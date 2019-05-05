package main
import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"bytes"
)

func TestIndex(t *testing.T) {
	expected := "{\"msg\":\"Hello, World!\"}"
	req, err := http.NewRequest("GET", "localhost:8080/", nil)

	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	Index(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	// do something
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}
	if msg := string(bytes.TrimSpace(b)); msg != expected {
		t.Errorf("expected message %q; got %q", expected, msg)
	}
}

func TestAnother(t *testing.T) {
	expected := "{\"msg\":\"Another Hello, World!\"}"
	req, err := http.NewRequest("GET", "localhost:8080/another", nil)

	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	Another(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	// do something
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}
	if msg := string(bytes.TrimSpace(b)); msg != expected {
		t.Errorf("expected message %q; got %q", expected, msg)
	}
}