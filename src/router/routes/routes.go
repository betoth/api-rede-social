package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes represents all API routes
type Routes struct {
	URI                    string
	Method                 string
	Function               func(w http.ResponseWriter, r *http.Request)
	AuthenticationRequires bool
}

// SetUp ...
func SetUp(r *mux.Router) *mux.Router {

	routers := userRoutes

	for _, router := range routers {

		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}

	return r
}
