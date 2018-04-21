package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("Update visit attendance by URL", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/favicon.ico", strings.NewReader(""))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to fetch favicon:", response.Code)
		}
	})
}
