package router

import (
	"api-rede-social/src/router/routes"

	"github.com/gorilla/mux"
)

// Create generates and returns a router with configured routes
func Create() *mux.Router {
	r := mux.NewRouter()
	return routes.SetUp(r)
}
