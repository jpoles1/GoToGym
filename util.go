package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/subosito/gotenv"
)

//envUsingMongoAtlas denotes the usage of a MongoDB atlas formatted URI.
var envProduction, envUsingMongoAtlas bool
var envBindPort, envBindIP, envBindURL, envMongoURI, envSMTPURI, envSMTPSender, envSMTPPass string
var envSMPTPPort int

func loadBoolEnv(varName string) bool {
	if os.Getenv(varName) == "" {
		color.Yellow(fmt.Sprintf("Missing %s value in .env file, automatically setting to false.\nSet a boolean value for %s in your .env file to disable this warning.", varName, varName))
		return false
	}
	usingMongoAtlas, err := strconv.ParseBool(os.Getenv(varName))
	if err != nil {
		color.Yellow(fmt.Sprintf("%s value must be a valid bool (true or false)\n Automatically setting to false.", varName))
		return false
	}
	return usingMongoAtlas

}

func loadEnv() {
	//Load Env
	gotenv.Load()

	//Setup global env variables
	envProduction = loadBoolEnv("PRODUCTION")
	envUsingMongoAtlas = loadBoolEnv("MONGO_ATLAS")

	if os.Getenv("BIND_PORT") == "" {
		log.Fatal("Missing BIND_PORT value in .env file.")
	}
	envBindPort = os.Getenv("BIND_PORT")

	if os.Getenv("BIND_IP") == "" {
		log.Fatal("Missing BIND_IP value in .env file.")
	}
	envBindIP = os.Getenv("BIND_IP")

	if os.Getenv("BIND_URL") == "" {
		log.Fatal("Missing BIND_URL value in .env file.")
	}
	envBindURL = os.Getenv("BIND_URL")

	if os.Getenv("MONGO_URI") == "" {
		log.Fatal("Missing MONGO_URI value in .env file.")
	}
	envMongoURI = os.Getenv("MONGO_URI")

	if os.Getenv("SMTP_SERVER") == "" {
		log.Fatal("Missing SMTP_SERVER value in .env file.")
	}
	envSMTPURI = os.Getenv("SMTP_SERVER")

	if port, found := os.LookupEnv("SMTP_PORT"); !found {
		log.Fatal("Missing SMTP_PORT value in .env file.")
	} else {
		var err error
		if envSMPTPPort, err = strconv.Atoi(port); err != nil {
			log.Fatal("SMTP_PORT value must be a number.")
		}
	}

	if os.Getenv("SMTP_SENDER") == "" {
		log.Fatal("Missing SMTP_SENDER value in .env file.")
	}
	envSMTPSender = os.Getenv("SMTP_SENDER")

	if os.Getenv("SMTP_PASS") == "" {
		log.Fatal("Missing SMTP_PASS value in .env file.")
	}
	envSMTPPass = os.Getenv("SMTP_PASS")
}

/*func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}*/

func errCheck(taskDescription string, err error) {
	if err != nil {
		log.Println("Error w/ " + taskDescription + ": " + err.Error())
	}
}

func sendAlert(doEmail bool, subject string, alertText string) {
	color.Red("High Importance Alert:")
	log.Println(alertText)
	//TODO add an email alert for high importance alerts (like network failures)
	if doEmail {
		err := sendEmail(envSMTPSender, "Alert:"+subject, "Alert:<br>"+alertText)
		errCheck("Sending alert email", err)
	}
}
