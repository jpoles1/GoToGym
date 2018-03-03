package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/globalsign/mgo/bson"
)

var testRouter = initRouter()

func TestGymVisitHandler(t *testing.T) {
	userID := bson.NewObjectId()
	userSecret := "secret"
	gymVisitID := bson.NewObjectId()
	t.Run("Create User by URL", func(t *testing.T) {
		createUserDocument(UserDocument{
			userID,
			userSecret,
			"jpdev.noreply@gmail.com",
			"Jordan", "Poles",
			false, []byte{},
		}, "password")
	})
	t.Run("Create visit by URL", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/api/gymvisit", strings.NewReader("{\"apikey\": \"secret\", \"title\": \"URL Test\", \"desc\": \"Test Description\", \"startTime\": \"\" , \"endTime\": \"\" }}"))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to submit to gymvisit endpoint. Err code:", response.Code)
		}
	})
	t.Run("Create visit", func(t *testing.T) {
		createGymVisitDocument(&GymVisitDocument{
			gymVisitID,
			userID,
			"Function Test",
			"Description Test",
			"", "",
			AttendanceUnset,
		})
	})
	t.Run("Update visit attendance by URL", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/verifyvisit/"+gymVisitID.Hex()+"/"+userSecret+"/yes", nil)
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to submit to verifyvisit endpoint. Err code:", response.Code, response.Body)
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
	t.Run("Delete All GymVisits", func(t *testing.T) {
		//deleteAllGymVisitDocuments()
	})
}
