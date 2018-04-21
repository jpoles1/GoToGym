package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/globalsign/mgo/bson"
)

var testRouter = initRouter()

func TestValidateTemplates(t *testing.T) {
	validateTemplates()
}
func TestServeErrorCode(t *testing.T) {
	serveErrorCode(httptest.NewRecorder(), http.StatusForbidden, "test")
}
func TestGymVisitHandler(t *testing.T) {
	userID := bson.NewObjectId()
	apiKey := "secret"
	email := "jpdev.noreply@gmail.com"
	password := "testpass"
	gymVisitID := bson.NewObjectId()
	t.Run("Create User by URL", func(t *testing.T) {
		createUserDocument(UserDocument{
			userID,
			apiKey,
			email,
			"Jordan", "Poles",
			false, []byte{},
		}, password)
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
			"March 3, 2018 at 09:30AM", "March 3, 2018 at 09:50AM",
			AttendanceUnset,
		})
	})
	t.Run("Update visit attendance by URL", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/verifyvisit/"+gymVisitID.Hex()+"/"+apiKey+"/yes", nil)
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to submit to verifyvisit endpoint. Err code:", response.Code, response.Body)
		}
	})
	t.Run("Login to account", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/api/login", strings.NewReader("{\"email\": \""+email+"\", \"password\": \""+password+"\"}"))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to login. Err code:", response.Code, response.Body)
		}
		if response.Body.String() == "" {
			t.Error("Incorrect login credentials!")
		}
	})
	t.Run("Fetch visit list", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/visitlist/"+apiKey, nil)
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to fetch from visitlist endpoint. Err code:", response.Code, response.Body)
		}
		fmt.Println(response.Body)
	})
	t.Run("Reset password", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/api/resetpassword/"+email+"/"+apiKey, nil)
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to reset password. Err code:", response.Code, response.Body)
		}
		var err error
		password, err = resetUserPassword(userID, email)
		errCheck("Resetting user password during test", err)
	})
	t.Run("Update password", func(t *testing.T) {
		newPassword := "testpass"
		request, _ := http.NewRequest("POST", "/api/updatepassword", strings.NewReader("{\"email\": \""+email+"\", \"oldPassword\": \""+password+"\", \"newPassword\": \""+newPassword+"\"}"))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to update password. Err code:", response.Code, response.Body)
		}
		if response.Body.String() == "" {
			t.Error("Incorrect login credentials!")
		}
		password = newPassword
	})
	t.Run("Login to account to check update", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/api/login", strings.NewReader("{\"email\": \""+email+"\", \"password\": \""+password+"\"}"))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to login. Err code:", response.Code, response.Body)
		}
		if response.Body.String() == "" {
			t.Error("Incorrect login credentials!")
		}
	})
	t.Run("Delete User", func(t *testing.T) {
		deleteUserDocument(userID)
	})
}

func TestUserRegistration(t *testing.T) {
	t.Run("Create User Via API", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/api/registration", strings.NewReader("{\"email\": \"jpdev.noreply@gmail.com\", \"firstName\": \"Jordan\", \"lastName\": \"Poles\", \"password\": \"password\"}"))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to submit to registration endpoint. Err code:", response.Code, response.Body)
		}
	})
	t.Run("Login to account", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/api/login", strings.NewReader("{\"email\": \"jpdev.noreply@gmail.com\", \"password\": \"password\"}"))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to login. Err code:", response.Code, response.Body)
		}
		if response.Body.String() == "" {
			t.Error("Incorrect login credentials!")
		}
	})
	t.Run("Create Second User Via API", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/api/registration", strings.NewReader("{\"email\": \"jpdev.noreply1@gmail.com\", \"firstName\": \"Jordan\", \"lastName\": \"Poles\", \"password\": \"password2\"}"))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to submit to registration endpoint. Err code:", response.Code, response.Body)
		}
	})
	t.Run("Login to Second Account", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/api/login", strings.NewReader("{\"email\": \"jpdev.noreply1@gmail.com\", \"password\": \"password2\"}"))
		response := httptest.NewRecorder()
		testRouter.ServeHTTP(response, request)
		if response.Code != 200 {
			t.Error("Failed to login. Err code:", response.Code, response.Body)
		}
		if response.Body.String() == "" {
			t.Error("Incorrect login credentials!")
		}
	})
	t.Run("Delete All Users", func(t *testing.T) {
		deleteAllUserDocuments()
	})
	t.Run("Delete All GymVisits", func(t *testing.T) {
		deleteAllGymVisitDocuments()
	})
}
