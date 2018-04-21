package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"time"

	mgo "github.com/globalsign/mgo"
)

var mongoConn *mgo.Session

func mongoConnect(mongoURI string) (*mgo.Session, error) {
	if envUsingMongoAtlas {
		return mongoAtlasConnect(mongoURI)
	}
	return mgo.Dial(envMongoURI)
}

func mongoAtlasConnect(mongoURI string) (*mgo.Session, error) {
	//URI without ssl=true
	mongoURI = strings.Replace(mongoURI, "ssl=true", "", 1)
	dialInfo, err := mgo.ParseURL(mongoURI)
	if err != nil {
		return &mgo.Session{}, err
	}
	//Below part is similar to above.
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	return mgo.DialWithInfo(dialInfo)
}

func dbLoad() *mgo.Session {
	var err error
	var mongoDB *mgo.Session
	attemptNum := 1
	maxWaitTime := time.Duration(30) * time.Second
	//Set initial time to wait on retry
	for err != nil || mongoDB == nil {
		mongoDB, err = mongoConnect(envMongoURI)
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
