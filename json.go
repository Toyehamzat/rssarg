package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responsdWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	dat, err := json.Marshal(payload)

	if err != nil {
		log.Printf("failed to marshal json response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
