package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Error("Failed to fetch homepage")
	}
}

func TestVisitViewHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/visitview/apikey", nil)
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Error("Failed to fetch visit view page")
	}
}
func TestLoginPageHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/login", nil)
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Error("Failed to fetch login page")
	}
}
func TestRegistrationPageHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/registration", nil)
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Error("Failed to fetch registration page")
	}
}
