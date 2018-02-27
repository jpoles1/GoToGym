package main

import (
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/subosito/gotenv"
)

var envBindPort, envBindIP, envBindURL, envMongoURI, envSMTPURI, envSMTPSender, envSMTPPass string
var envSMPTPPort int

func loadEnv() {
	//Load Env
	gotenv.Load()

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

func errCheck(taskDescription string, err error) {
	if err != nil {
		log.Println("Error w/ " + taskDescription + ": " + err.Error())
	}
}

func sendAlert(email bool, subject string, alertText string) {
	color.Red("High Importance Alert:")
	log.Println(alertText)
	//TODO add an email alert for high importance alerts (like network failures)
	if email {
		err := sendEmail(envSMTPSender, "Alert:"+subject, "Alert:<br>"+alertText)
		errCheck("Sending alert email", err)
	}
}
