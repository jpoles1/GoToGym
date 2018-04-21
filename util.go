package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	raven "github.com/getsentry/raven-go"
	"github.com/subosito/gotenv"
)

//envUsingMongoAtlas denotes the usage of a MongoDB atlas formatted URI.
var envProduction, envUsingMongoAtlas bool
var envBindPort, envBindIP, envBindURL, envMongoURI, envSMTPURI, envSMTPSender, envSMTPPass, envSentryDSN string
var envSMTPPort int

func loadBoolEnv(varName string) bool {
	if os.Getenv(varName) == "" {
		color.Yellow(fmt.Sprintf("Missing %s value in .env file, automatically setting to false.\nSet a boolean value for %s in your .env file to disable this warning.", varName, varName))
		return false
	}
	boolEnv, err := strconv.ParseBool(os.Getenv(varName))
	if err != nil {
		color.Yellow(fmt.Sprintf("%s value must be a valid bool (true or false)\n Automatically setting to false.", varName))
		return false
	}
	return boolEnv

}
func loadStringEnv(varName string) string {
	if os.Getenv(varName) == "" {
		log.Fatal(fmt.Sprintf("Missing %s value in .env file.", varName))
	}
	return os.Getenv(varName)
}

func loadIntEnv(varName string) int {
	stringEnv, found := os.LookupEnv(varName)
	if !found {
		log.Fatal(fmt.Sprintf("Missing %s value in .env file.", varName))
	}
	intEnv, err := strconv.Atoi(stringEnv)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s value must be an integer.", varName))
	}
	return intEnv
}

func loadEnv() {
	//Load Env
	gotenv.Load()

	//Setup global env variables
	envProduction = loadBoolEnv("PRODUCTION")
	envUsingMongoAtlas = loadBoolEnv("MONGO_ATLAS")

	envBindPort = loadStringEnv("BIND_PORT")
	envBindIP = loadStringEnv("BIND_IP")
	envBindURL = loadStringEnv("BIND_URL")
	envMongoURI = loadStringEnv("MONGO_URI")
	envSMTPURI = loadStringEnv("SMTP_SERVER")
	envSMTPPort = loadIntEnv("SMTP_PORT")
	envSMTPSender = loadStringEnv("SMTP_SENDER")
	envSMTPPass = loadStringEnv("SMTP_PASS")
	envSentryDSN = loadStringEnv("SENTRY_DSN")
}

/*func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}*/
func loadSentry() {
	raven.SetDSN(envSentryDSN)
}

func errCheck(taskDescription string, err error) {
	if err != nil {
		raven.CaptureError(err, map[string]string{"msg": taskDescription})
		log.Println("Error w/ " + taskDescription + ": " + err.Error())
	}
}

func sendAlert(doEmail bool, subject string, alertText string) {
	color.Red("High Importance Alert:")
	raven.CaptureError(errors.New(subject), map[string]string{"level": "Alert", "msg": alertText})
	log.Println(alertText)
	//TODO add an email alert for high importance alerts (like network failures)
	if doEmail {
		err := sendEmail(envSMTPSender, "Alert:"+subject, "Alert:<br>"+alertText)
		errCheck("Sending alert email", err)
	}
}
