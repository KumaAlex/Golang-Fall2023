package main

import (
	"ProjectGO/internal/data"
	"ProjectGO/internal/validator"
	"errors"
	"fmt"
	"net/http"
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

	if data.ValidateLaptopBag(v, laptopBag); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.LaptopBags.Insert(laptopBag)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/laptopBags/%d", laptopBag.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"laptopBag": laptopBag}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showLaptopBagHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	laptopBag, err := app.models.LaptopBags.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrBagNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"laptopBag": laptopBag}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) updateLaptopBagHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	laptopBag, err := app.models.LaptopBags.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrBagNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Brand      *string      `json:"brand"`
		Color      *string      `json:"color"`
		Weight     *data.Weight `json:"weight"`
		Dimensions []float32    `json:"dimensions"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Brand != nil {
		laptopBag.Brand = *input.Brand
	}
	if input.Color != nil {
		laptopBag.Color = *input.Color
	}
	if input.Weight != nil {
		laptopBag.Weight = *input.Weight
	}
	if input.Dimensions != nil {
		laptopBag.Dimensions = input.Dimensions
	}

	v := validator.New()
	if data.ValidateLaptopBag(v, laptopBag); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.LaptopBags.Update(laptopBag)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"laptopBag": laptopBag}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteLaptopBagHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.LaptopBags.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrBagNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "laptop bag successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listLaptopBagsHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Brand string
		Color string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Brand = app.readString(qs, "brand", "")
	input.Color = app.readString(qs, "color", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{"id", "brand", "color", "model", "material", "compartments", "-id", "-brand", "-color", "-model", "-material", "-compartments"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	laptopBags, metadata, err := app.models.LaptopBags.GetAll(input.Brand, input.Color, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"laptopBags": laptopBags, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
