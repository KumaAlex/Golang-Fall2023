package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/laptopBags", app.listLaptopBagsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/laptopBags", app.createLaptopBagHandler)
	router.HandlerFunc(http.MethodGet, "/v1/laptopBags/:id", app.showLaptopBagHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/laptopBags/:id", app.updateLaptopBagHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/laptopBags/:id", app.deleteLaptopBagHandler)

	return router
}

//BODY='{"brand":"Urban lifestyle","color":"black","weight":"1.6 kg", "Dimensions":[10,16,41.6]}'
//curl -i -d "$BODY" localhost:4000/v1/laptopBags
//BODY='{"brand":"Pandec","color":"black and white","weight":"4 kg", "Dimensions":[40, 40, 40]}'
//curl -i -d "$BODY" localhost:4000/v1/laptopBags
//BODY='{"brand":"Nig","color":"black","weight":"0.1 kg", "Dimensions":[4,5,6]}'
//curl -i -d "$BODY" localhost:4000/v1/laptopBags

//BODY='{"brand":"Urban lifestyle","color":"black","weight":"69 kg", "Dimensions":[10,16,41.6]}'
//curl -X PUT -d "$BODY" localhost:4000/v1/laptopBags/1

//curl localhost:4000/v1/laptopBags/1

//curl -X DELETE localhost:4000/v1/laptopBags/3
