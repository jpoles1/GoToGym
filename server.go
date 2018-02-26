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
	//Define Routes
	router.HandleFunc("/", homePageHandler)
	router.HandleFunc("/api/gymvisit", apiHandlers["gymvisit"]).Methods("POST")
	//Serve static files stored in html/static
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	return router
}
