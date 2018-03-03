package main

import (
	"errors"

	"github.com/globalsign/mgo/bson"
)

//Attendance is a type used to store the info on a user's attendance at a given gym calendar event
type Attendance int8

const (
	//AttendanceUnset is the default value before user response
	AttendanceUnset Attendance = -1
	//AttendanceMissed is the value for a user who did not go to the scheduled gym visit
	AttendanceMissed Attendance = 0
	//AttendanceAttended is the value for a user who did go to the scheduled gym visit
	AttendanceAttended Attendance = 1
)

//GymVisitDocument is a stucture to contain an entry in the GymVisit mgo collection
type GymVisitDocument struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	UserID      bson.ObjectId
	Title       string
	Description string
	StartTime   string
	EndTime     string
	Attended    Attendance
}

func createGymVisitDocument(doc *GymVisitDocument) error {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	err := mongoSesh.DB("gotogym").C("gymvisits").Insert(doc)
	errCheck("Inserting gym visit into DB", err)
	return err
}

func findGymVisitDocumentByID(documentID bson.ObjectId) (*GymVisitDocument, error) {
	var err error
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	searchParams := bson.M{
		"_id": documentID,
	}
	var gymVisitDocument GymVisitDocument
	mongoSesh.DB("gotogym").C("gymvisits").Find(searchParams).One(&gymVisitDocument)
	if gymVisitDocument.ID == bson.ObjectId("") {
		err = errors.New("Could not fetch a document with this ID")
	}
	return &gymVisitDocument, err
}
func updateGymVisitDocumentByID(doc *GymVisitDocument) error {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	_, err := mongoSesh.DB("gotogym").C("gymvisits").Upsert(bson.M{"_id": doc.ID}, doc)
	return err
}

func deleteAllGymVisitDocuments() {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	_, err := mongoSesh.DB("gotogym").C("gymvisits").RemoveAll(bson.M{})
	errCheck("Removing all gymvisits from DB", err)
}
