package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes create routes to router
type Routes struct {
	URI         string
	Method      string
	Handler     func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// ConfigRoutes Config routes in Router
func ConfigRoutes(r *mux.Router) *mux.Router {
	for _, route := range userRoutes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return r
}
