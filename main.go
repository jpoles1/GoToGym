package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func main() {
	loadEnv()
	//Add Microservice functions to this list
	microservices := []func(){
		func() { webServer() },
	}

	//Create a wait group for all microservices
	var wg sync.WaitGroup
	wg.Add(len(microservices))

	//create a thread for each microservice
	for _, microservice := range microservices {
		localmicroservice := microservice
		go func() {
			defer wg.Done()
			localmicroservice()
		}()
	}

	//wait for all microservices to complete
	wg.Wait()
}

func webServer() {
	validateTemplates()
	//Spool up http server
	color.Green("Starting Web server on port: %s", envBindPort)
	color.Green("Access the web server at: http://%s:%s", envBindIP, envBindPort)
	logrus.Fatal(http.ListenAndServe(envBindIP+":"+envBindPort, initRouter()))
	fmt.Println("Terminating TransitSign Web Server...")
}
