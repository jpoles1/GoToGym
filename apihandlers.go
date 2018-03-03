package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

var apiHandlers = apiHandlerSetup()

func apiHandlerSetup() map[string]func(http.ResponseWriter, *http.Request) {
	var apiHandlers = map[string]func(http.ResponseWriter, *http.Request){}
	apiHandlers["gymvisit"] = func(w http.ResponseWriter, r *http.Request) {
		type apiStruct struct {
			APIKey      string `json:"apikey"`
			Title       string `json:"title"`
			Description string `json:"desc"`
			StartTime   string `json:"startTime"`
			EndTime     string `json:"endTime"`
		}
		decoder := json.NewDecoder(r.Body)
		var apiData apiStruct
		err := decoder.Decode(&apiData)
		errCheck("Decoding gymvisit API request", err)
		defer r.Body.Close()
		log.Println(apiData)
		//Check that all required fields are filled
		if apiData.APIKey == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Need to fill all fields (apikey)"))
			return
		}
		userData, err := findUserDocumentByAPIKey(apiData.APIKey)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("API key not found"))
			return
		}
		visitData := GymVisitDocument{
			bson.NewObjectId(),
			(*userData).ID,
			apiData.Title,
			apiData.Description,
			apiData.StartTime,
			apiData.EndTime,
			AttendanceUnset,
		}
		err = createGymVisitDocument(&visitData)
		errCheck("Decoding gymvisit API request", err)
		err = sendGymVisitCheckinEmail(visitData, userData)
		errCheck("Sending gymvisit email", err)
		w.Write([]byte("Gym Visit Entry Received"))
	}
	apiHandlers["newuser"] = func(w http.ResponseWriter, r *http.Request) {
		type apiStruct struct {
			Email     string `json:"email"`
			FirstName string `json:"firstname"`
			LastName  string `json:"lastname"`
			Password  string `json:"password"`
		}
		decoder := json.NewDecoder(r.Body)
		var apiData apiStruct
		err := decoder.Decode(&apiData)
		errCheck("Decoding newuser API request", err)
		defer r.Body.Close()
		log.Println(apiData)
		if apiData.Email == "" || apiData.FirstName == "" || apiData.LastName == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Need to fill all fields (email, firstname, lastname)"))
			return
		}
		userCount := findUserDocumentByEmail(apiData.Email)
		if userCount > 0 {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Email already in use"))
			return
		}
		apiKey := uuid.NewV4().String()
		newUserData := UserDocument{
			bson.NewObjectId(),
			apiKey,
			apiData.Email,
			apiData.FirstName, apiData.LastName,
			false, []byte{},
		}
		err = sendRegistrationEmail(&newUserData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to send email: " + err.Error()))
			return
		}
		createUserDocument(newUserData, apiData.Password)
		w.Write([]byte("New User Entry Received"))
	}
	apiHandlers["verifyvisit"] = func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		documentID := vars["documentID"]
		apiKey := vars["apiKey"]
		visitResponse := vars["response"]
		gymVisitData, err := findGymVisitDocumentByID(bson.ObjectIdHex(documentID))
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Document not found:" + err.Error()))
			return
		}
		userData, err := findUserDocumentByAPIKey(apiKey)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("API key not found"))
			return
		}
		if gymVisitData.UserID == userData.ID {
			if visitResponse == "yes" {
				gymVisitData.Attended = AttendanceAttended
			} else if visitResponse == "no" {
				gymVisitData.Attended = AttendanceMissed
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid response"))
				return
			}
			err := updateGymVisitDocumentByID(gymVisitData)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Failed to update DB entry: " + err.Error()))
				return
			}
			w.Write([]byte("Successfully updated visit attendance!"))
		}
	}
	return apiHandlers
}
