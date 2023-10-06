package main

import (
	"ProjectGO/internal/data"
	"ProjectGO/internal/validator"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createLaptopBagHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Brand      string      `json:"brand"`
		Color      string      `json:"color"`
		Weight     data.Weight `json:"weight"`
		Dimensions []float32   `json:"dimensions"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	laptopBag := &data.LaptopBag{
		Brand:      input.Brand,
		Color:      input.Color,
		Weight:     input.Weight,
		Dimensions: input.Dimensions,
	}

	v := validator.New()

	// Call the ValidateMovie() function and return a response containing the errors if
	// any of the checks fail.
	if data.ValidateMovie(v, laptopBag); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Dump the contents of the input struct in a HTTP response.
	fmt.Fprintf(w, "%+v\n", input)
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
