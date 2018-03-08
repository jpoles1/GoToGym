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
func visitViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apiKey := vars["apiKey"]
	t := template.New("visitview")
	t, _ = template.ParseFiles("templates/visitview.gohtml")
	listViewData := struct {
		APIKey string
	}{
		apiKey,
	}
	t.Execute(w, listViewData)
}
func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("login")
	t, _ = template.ParseFiles("templates/login.gohtml")
	t.Execute(w, nil)
}
func registrationPageHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("registration")
	t, _ = template.ParseFiles("templates/registration.gohtml")
	t.Execute(w, nil)
}
