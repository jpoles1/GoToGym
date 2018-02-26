package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	loadEnv()
	webServer()
}

func webServer() {
	if _, err := template.ParseGlob("templates/*.gohtml"); err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	//Define Routes

	//Serve static files stored in html/static
	router.HandleFunc("/", homePageHandler)
	router.HandleFunc("/api/gymvisit", apiHandlers["gymvisit"]).Methods("POST")
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	//Spool up http server
	color.Green("Starting Web server on port: %s", envBindPort)
	color.Green("Access the web server at: http://%s:%s", envBindIP, envBindPort)
	logrus.Fatal(http.ListenAndServe(envBindIP+":"+envBindPort, router))
	fmt.Println("Terminating TransitSign Web Server...")
}
