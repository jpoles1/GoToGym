package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("home")
	t, _ = template.ParseFiles("templates/index.gohtml")
	t.Execute(w, nil)
}
func visitListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apiKey := vars["apiKey"]
	t := template.New("visitlist")
	t, _ = template.ParseFiles("templates/visitview.gohtml")
	listViewData := struct {
		APIKey string
	}{
		apiKey,
	}
	t.Execute(w, listViewData)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("visitlist")
	t, _ = template.ParseFiles("templates/login.gohtml")
	t.Execute(w, nil)
}
