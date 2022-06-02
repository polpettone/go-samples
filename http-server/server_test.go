package main

import (
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost", nil)
	w := httptest.NewRecorder()

	echo(w, req)

	result := w.Result()

	if result.StatusCode != 200 {
		t.Errorf("wanted %d got %d", 200, result.StatusCode)
	}

}
