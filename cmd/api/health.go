package main

import "net/http"

// healthCheck godoc
//
// @Summary check the API health
// @Descripion check the API health
// @ID health
// @Router /health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Health OK!"))
}
