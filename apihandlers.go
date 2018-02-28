package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
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
		userData, err := findUserDocumentByAPIKey(apiData.APIKey)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
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
		err = sendGymVisitCheckin(visitData, userData)
		errCheck("Sending gymvisit email", err)
		w.Write([]byte("Gym Visit Entry Received"))
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
