package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `jason:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Indexing backend up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) Search(w http.ResponseWriter, r *http.Request) {

}
