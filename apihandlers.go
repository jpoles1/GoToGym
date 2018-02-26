package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var apiHandlers = apiHandlerSetup()

func apiHandlerSetup() map[string]func(http.ResponseWriter, *http.Request) {
	var apiHandlers = map[string]func(http.ResponseWriter, *http.Request){}
	apiHandlers["gymvisit"] = func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t interface{}
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		log.Println(t)
		w.Write([]byte("TEST"))
	}
	fmt.Println("PREPPED API ROUTER")
	return apiHandlers
}
