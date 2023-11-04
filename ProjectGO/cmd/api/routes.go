package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/laptopBags", app.listLaptopBagsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/laptopBags", app.createLaptopBagHandler)
	router.HandlerFunc(http.MethodGet, "/v1/laptopBags/:id", app.showLaptopBagHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/laptopBags/:id", app.updateLaptopBagHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/laptopBags/:id", app.deleteLaptopBagHandler)

	return app.recoverPanic(app.rateLimit(router))

}
