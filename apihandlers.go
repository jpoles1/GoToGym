package main

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

var apiHandlers = apiHandlerSetup()

func apiHandlerSetup() map[string]func(http.ResponseWriter, *http.Request) {
	var apiHandlers = map[string]func(http.ResponseWriter, *http.Request){}
	//Use this endpoint to get a list of gym visits from a given user's account
	apiHandlers["login"] = func(w http.ResponseWriter, r *http.Request) {
		type apiStruct struct {
			Email             string `json:"email"`
			PlaintextPassword string `json:"password"`
		}
		decoder := json.NewDecoder(r.Body)
		var apiData apiStruct
		err := decoder.Decode(&apiData)
		errCheck("Decoding login API request", err)
		defer r.Body.Close()
		userData, err := checkUserCredentials(apiData.Email, apiData.PlaintextPassword)
		if err != nil {
			w.Write([]byte(""))
			return
		}
		w.Write([]byte(userData.APIKey))
	}
	//Use this endpoint to get a list of gym visits from a given user's account
	apiHandlers["visitlist"] = func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		apiKey := vars["apiKey"]
		userDoc, err := findUserDocumentByAPIKey(apiKey)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("API key not found"))
			return
		}
		visitDocs := findGymVisitDocumentsByUserID(userDoc.ID)
		jsonData, err := json.Marshal(visitDocs)
		errCheck("Endoding visit list JSON", err)
		w.Write(jsonData)
	}
	//Use this endpoint to add a new gym visit to a given user's account
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
	//Use this endpoint to add a new user account
	apiHandlers["registration"] = func(w http.ResponseWriter, r *http.Request) {
		validateEmail := func(email string) bool {
			Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
			return Re.MatchString(email)
		}
		type apiStruct struct {
			Email     string `json:"email"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Password  string `json:"password"`
		}
		decoder := json.NewDecoder(r.Body)
		var apiData apiStruct
		err := decoder.Decode(&apiData)
		errCheck("Decoding newuser API request", err)
		defer r.Body.Close()
		if apiData.Email == "" || apiData.FirstName == "" || apiData.LastName == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Need to fill all fields"))
			return
		}
		if !validateEmail(apiData.Email) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Invalid email address."))
			return
		}
		userCount := findUserDocumentByEmail(apiData.Email)
		if userCount > 0 {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Email in use"))
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
		w.Write([]byte(newUserData.APIKey))
	}
	//Use this endpoint to update a gym visit with user attendance
	apiHandlers["verifyvisit"] = func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		documentID := vars["documentID"]
		apiKey := vars["apiKey"]
		visitResponse := vars["response"]
		gymVisitData, err := findGymVisitDocumentByID(bson.ObjectIdHex(documentID))
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Document not found: " + err.Error()))
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
				gymVisitData.Attendance = AttendanceAttended
			} else if visitResponse == "no" {
				gymVisitData.Attendance = AttendanceMissed
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
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(`
				<meta http-equiv="refresh" content="1.5;url=/visitview/` + apiKey + `"/>
				Successfully updated visit attendance! <a href="/visitview/` + apiKey + `">Redirecting...</a>
			`))
		}
	}
	return apiHandlers
}
