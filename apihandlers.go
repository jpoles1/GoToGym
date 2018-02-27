package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

//GymVisitDocument is a stucture to contain an entry in the GymVisit mgo collection
type GymVisitDocument struct {
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
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		log.Println(apiData)
		userID := findUserDocumentByAPIKey(apiData.APIKey)
		if userID == bson.ObjectId("") {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		storeGymVisit(GymVisitDocument{
			userID,
			apiData.Title,
			apiData.Description,
			apiData.StartTime,
			apiData.EndTime,
		})
		w.Write([]byte("Gym Visit Entry Received"))
	}
	return apiHandlers
}
func storeGymVisit(doc GymVisitDocument) {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	err := mongoSesh.DB("gotogym").C("gymvisits").Insert(doc)
	errCheck("Inserting gym visit into DB", err)
}
