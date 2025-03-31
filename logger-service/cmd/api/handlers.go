package main

import (
	"log"
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {

	// Read JSON into var
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)
	log.Println("Reached Logger-service WriteLog")

	// insert Data

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		log.Println("Error While inserthing data into Log: ", w, err)
		return
	}
	resp := jsonResponse{
		Error:   false,
		Message: "Logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)

}
