package main

import (
	"net/http"
)

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	responsdWithJSON(w, 200, struct{}{})
}
