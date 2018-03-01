package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	uuid "github.com/satori/go.uuid"
)

//GymVisitDocument is a stucture to contain an entry in the GymVisit mgo collection
type GymVisitDocument struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	UserID      bson.ObjectId
	Title       string
	Description string
	StartTime   string
	EndTime     string
}

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
			w.WriteHeader(http.StatusForbidden)
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
		}
		err = storeGymVisit(&visitData)
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
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Need to fill all fields (email, firstname, lastname)"))
			return
		}
		userCount := findUserDocumentByEmail(apiData.Email)
		if userCount > 0 {
			w.WriteHeader(http.StatusBadRequest)
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
		createUserDocument(newUserData, apiData.Password)
		err = sendRegistrationEmail(&newUserData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to send email: " + err.Error()))
			return
		}
		w.Write([]byte("New User Entry Received"))
	}
	return apiHandlers
}
func storeGymVisit(doc *GymVisitDocument) error {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	err := mongoSesh.DB("gotogym").C("gymvisits").Insert(doc)
	errCheck("Inserting gym visit into DB", err)
	return err
}
