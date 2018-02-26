package main

import "github.com/globalsign/mgo/bson"

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

func createUserDocument(userDoc UserDocument) {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	err := mongoSesh.DB("transitserver").C("users").Insert(userDoc)
	errCheck("Inserting user into DB", err)
}
