package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var apiHandlers = map[string]func(http.ResponseWriter, *http.Request){}

func init() {
	apiHandlers["gymvisit"] = func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t interface{}
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		log.Println(t)
	}
}
