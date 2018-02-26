package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	fmt.Println(response)
	if response.Code != 200 {
		t.Error("Failed to fetch homepage")
	}
}
