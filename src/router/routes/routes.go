package routes

import (
	"api-rede-social/src/middlewares"
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
	routers = append(routers, loginRoutes)
	routers = append(routers, publicationRoutes...)

	for _, router := range routers {
		if router.AuthenticationRequires {
			r.HandleFunc(router.URI,
				middlewares.Logger(middlewares.Authentication(router.Function)),
			).Methods(router.Method)

		} else {
			r.HandleFunc(router.URI, middlewares.Logger(router.Function)).Methods(router.Method)

		}

	}

	return r
}
