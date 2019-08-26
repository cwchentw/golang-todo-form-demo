package main

import (
	"encoding/json"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string `json:"message"`
	}{
		"No valid resource",
	}

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound) // HTTP 404.
	w.Write(json)
}
