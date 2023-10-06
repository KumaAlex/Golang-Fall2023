package main

import (
	"ProjectGO/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createLaptopBagHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new LaptopBag")
}

func (app *application) showLaptopBagHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	laptopBag := data.LaptopBag{
		ID:           id,
		CreatedAt:    time.Now(),
		Brand:        "Urban lifestyle",
		Model:        "Day Back",
		Color:        "yellow",
		Material:     "cloth",
		Compartments: 2,
		Weight:       0.230,
		Dimensions:   []float32{11.4, 5.1, 14.2},
	}
	// Encode the struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, envelope{"laptopBag": laptopBag}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
