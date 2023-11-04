package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func healthcheckHandlerEncoder(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": "development",
		"version":     "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func healthcheckHandlerMarshal(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": "development",
		"version":     "1.0.0",
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
