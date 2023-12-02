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

	router.HandlerFunc(http.MethodGet, "/v1/laptopBags", app.requirePermission("laptopBags:read", app.listLaptopBagsHandler))
	router.HandlerFunc(http.MethodPost, "/v1/laptopBags", app.requirePermission("laptopBags:write", app.createLaptopBagHandler))
	router.HandlerFunc(http.MethodGet, "/v1/laptopBags/:id", app.requirePermission("laptopBags:read", app.showLaptopBagHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/laptopBags/:id", app.requirePermission("laptopBags:write", app.updateLaptopBagHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/laptopBags/:id", app.requirePermission("laptopBags:write", app.deleteLaptopBagHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
