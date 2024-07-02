package middlewares

import (
	"api-rede-social/src/authentication"
	"api-rede-social/src/response"
	"log"
	"net/http"
)

// Logger add log to requests
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}

}

// Authentication verify user login in
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			response.ErrorJSON(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}

}
