package main

import (
	"testing"

	"github.com/globalsign/mgo/bson"
)

func TestCreateUserDocument(t *testing.T) {
	createUserDocument(UserDocument{
		bson.NewObjectId(),
		"secret",
		"jpoles1@gmail.com",
		"Jordan", "Poles",
		false, []byte{},
	})
}
