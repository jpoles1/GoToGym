package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("home")
	t, _ = template.ParseFiles("templates/index.gohtml")
	homePageData := struct {
		Name string
	}{
		"JP",
	}
	t.Execute(w, homePageData)
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
