package main

import (
	"html/template"
	"net/http"
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
