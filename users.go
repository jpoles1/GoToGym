package main

type User struct {
	Email          string `json:"email" bson:"email"`
	LoginUUID      string `json:"-" bson:"login_uuid"`
	FirstName      string `json:"first_name" bson:"first_name"`
	LastName       string `json:"last_name" bson:"last_name"`
	EmailConfirmed bool   `json:"-" bson:"email_confirmed"`
	EmailUUID      string `json:"-" bson:"email_uuid"`
	HashedPassword []byte `json:"-" bson:"hashed_password"`
}
