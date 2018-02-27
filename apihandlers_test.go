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
		})
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
