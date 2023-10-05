package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Add a createLaptopBagHandler for the "POST /v1/laptopBags" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createLaptopBagHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new LaptopBag")
}

// Add a showLaptopBagHandler for the "GET /v1/laptopBag/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showLaptopBagHandler(w http.ResponseWriter, r *http.Request) {
	// When httprouter is parsing a request, any interpolated URL parameters will be
	// stored in the request context. We can use the ParamsFromContext() function to
	// retrieve a slice containing these parameter names and values.
	params := httprouter.ParamsFromContext(r.Context())
	// We can then use the ByName() method to get the value of the "id" parameter from
	// the slice. In our project all LaptopBags will have a unique positive integer ID, but
	// the value returned by ByName() is always a string. So we try to convert it to a
	// base 10 integer (with a bit size of 64). If the parameter couldn't be converted,
	// or is less than 1, we know the ID is invalid so we use the http.NotFound()
	// function to return a 404 Not Found response.
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Otherwise, interpolate the LaptopBag ID in a placeholder response.
	fmt.Fprintf(w, "show the details of LaptopBag %d\n", id)
}
