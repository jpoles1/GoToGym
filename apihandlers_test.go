package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/globalsign/mgo/bson"
)

/*func testRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", stationList)
	return router
}*/

var testRouter = initRouter()

func TestGymVisitHandler(t *testing.T) {
	var userID bson.ObjectId
	t.Run("Create User", func(t *testing.T) {
		userID = createUserDocument(UserDocument{
			bson.NewObjectId(),
			"secret",
			"jpdev.noreply@gmail.com",
			"Jordan", "Poles",
			false, []byte{},
		}, "password")
	})
	t.Run("Add visit", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/api/gymvisit", strings.NewReader("{\"apikey\": \"secret\", \"title\": \"Test Title\", \"desc\": \"Test Description\", \"startTime\": \"\" , \"endTime\": \"\" }}"))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to submit to gymvisit endpoint. Err code:", response.Code)
		}
	})
	t.Run("Delete User", func(t *testing.T) {
		deleteUserDocument(userID)
	})
}

func TestUserRegistration(t *testing.T) {
	t.Run("Create User Via API", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/api/newuser", strings.NewReader("{\"email\": \"jpdev.noreply@gmail.com\", \"firstname\": \"Jordan\", \"lastname\": \"Poles\"}"))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to submit to newuser endpoint. Err code:", response.Code, response.Body)
		}
	})
	t.Run("Delete All Users", func(t *testing.T) {
		deleteAllUserDocuments()
	})
}
