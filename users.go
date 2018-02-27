package main

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

//UserDocument is a stucture to contain an entry in the users mgo collection
type UserDocument struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	APIKey         string        `bson:"apikey"`
	Email          string        `json:"email" bson:"email"`
	FirstName      string        `json:"first_name" bson:"first_name"`
	LastName       string        `json:"last_name" bson:"last_name"`
	EmailConfirmed bool          `json:"-" bson:"email_confirmed"`
	HashedPassword []byte        `json:"-" bson:"hashed_password"`
}

func createUserDocument(userDoc UserDocument) bson.ObjectId {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	err := mongoSesh.DB("gotogym").C("users").Insert(userDoc)
	errCheck("Inserting user into DB", err)
	return userDoc.ID
}

func findUserDocumentByAPIKey(apiKey string) bson.ObjectId {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	searchParams := bson.M{
		"apikey": apiKey,
	}
	var userData UserDocument
	mongoSesh.DB("gotogym").C("users").Find(searchParams).One(&userData)
	fmt.Println()
	//TODO what happens when we don't get a response
	fmt.Println(apiKey, "UDATA", userData)
	return userData.ID
}

func deleteUserDocument(userID bson.ObjectId) {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	err := mongoSesh.DB("gotogym").C("users").RemoveId(userID)
	errCheck("Removing user from DB", err)
}
