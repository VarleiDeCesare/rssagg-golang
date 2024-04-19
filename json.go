package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to marshal json response %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Request failed with error 5XX: %s", message)
		w.WriteHeader(500)
		return
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, ErrorResponse{Error: message})
}
