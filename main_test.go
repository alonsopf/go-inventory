package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServer(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	HelloServer(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("HelloServer StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusOK)
	}
}

func TestJSONServer(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/json", nil)
	JSONServer(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("JSONServer StatusCode = %v, want %v", rr.Result().StatusCode, http.StatusOK)
	}
}
