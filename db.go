package main

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var mongoConn = dbLoad()

func dbLoad() *mgo.Session {
	var err error
	err = nil
	var mongoDB *mgo.Session
	attemptNum := 1
	maxWaitTime := time.Duration(30) * time.Second
	//Set initial time to wait on retry
	for err != nil || mongoDB == nil {
		mongoDB, err = mgo.Dial(envMongoURI)
		if err != nil || mongoDB == nil {
			//In milliseconds
			waitTime := time.Duration(2.5*float64(attemptNum)*1000) * time.Millisecond
			if waitTime > maxWaitTime {
				waitTime = maxWaitTime
			}
			sendAlert(attemptNum == 15, "Mongo Connection Failiure", fmt.Sprintf("Failed to connect to provided MongoDB URI (attepmt #%d; sleeping %v):\n%s", attemptNum, waitTime, err))
			time.Sleep(waitTime)
			attemptNum++
		}
	}
	return mongoDB
}
