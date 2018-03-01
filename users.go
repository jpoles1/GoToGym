package main

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

//UserDocument is a stucture to contain an entry in the users mgo collection
type UserDocument struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	APIKey         string        `json:"apikey" bson:"apikey"`
	Email          string        `json:"email" bson:"email"`
	FirstName      string        `json:"first_name" bson:"first_name"`
	LastName       string        `json:"last_name" bson:"last_name"`
	EmailConfirmed bool          `json:"-" bson:"email_confirmed"`
	PasswordHash   []byte        `json:"-" bson:"hashed_password"`
}

func createUserDocument(userDoc UserDocument, passString string) bson.ObjectId {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	// Hashing the password with the default cost of 10
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(passString), bcrypt.DefaultCost)
	errCheck("Encoding password to hash", err)
	userDoc.PasswordHash = passwordHash
	err = mongoSesh.DB("gotogym").C("users").Insert(userDoc)
	errCheck("Inserting user into DB", err)
	return userDoc.ID
}

func findUserDocumentByEmail(email string) int {
	var err error
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	searchParams := bson.M{
		"email": email,
	}
	userCount, err := mongoSesh.DB("gotogym").C("users").Find(searchParams).Count()
	errCheck("Finding User by Email", err)
	return userCount
}

func findUserDocumentByAPIKey(apiKey string) (*UserDocument, error) {
	var err error
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	searchParams := bson.M{
		"apikey": apiKey,
	}
	var userData UserDocument
	mongoSesh.DB("gotogym").C("users").Find(searchParams).One(&userData)
	if userData.ID == bson.ObjectId("") {
		err = errors.New("Could not fetch a user with this API key")
	}
	return &userData, err
}

func deleteUserDocument(userID bson.ObjectId) {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	err := mongoSesh.DB("gotogym").C("users").RemoveId(userID)
	errCheck("Removing user from DB", err)
}
func deleteAllUserDocuments() {
	mongoSesh := dbLoad()
	defer mongoSesh.Close()
	_, err := mongoSesh.DB("gotogym").C("users").RemoveAll(bson.M{})
	errCheck("Removing all users from DB", err)
}
