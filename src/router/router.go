package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate create a new Router with configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
