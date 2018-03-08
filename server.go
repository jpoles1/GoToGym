package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func validateTemplates() {
	if _, err := template.ParseGlob("templates/*.gohtml"); err != nil {
		panic(err)
	}
}

func initRouter() *mux.Router {
	router := mux.NewRouter()
	//Define Web Routes
	router.HandleFunc("/", homePageHandler)
	router.HandleFunc("/login", loginPageHandler)
	router.HandleFunc("/registration", registrationPageHandler)
	router.HandleFunc("/visitlist/{apiKey}", visitListHandler)
	//Define API Routes
	router.HandleFunc("/api/login", apiHandlers["login"]).Methods("POST")
	router.HandleFunc("/api/registration", apiHandlers["registration"]).Methods("POST")
	router.HandleFunc("/api/visitlist/{apiKey}", apiHandlers["visitlist"]).Methods("GET")
	router.HandleFunc("/api/gymvisit", apiHandlers["gymvisit"]).Methods("POST")
	router.HandleFunc("/api/verifyvisit/{documentID}/{apiKey}/{response}", apiHandlers["verifyvisit"]).Methods("GET")

	//Serve static files stored in html/static
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	return router
}
