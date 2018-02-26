package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*func testRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", stationList)
	return router
}*/

var testRouter = initRouter()

func TestGymVisitHandler(t *testing.T) {
	request, _ := http.NewRequest("POST", "/api/gymvisit", strings.NewReader("{\"apkikey\": \"secret\"}"))
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	fmt.Println(response)
	if response.Code != 200 {
		t.Error("Failed to fetch gymvisit endpoint")
	}
}
