package routes

import (
	"api-rede-social/src/controllers"
	"net/http"
)

var loginRoutes = Routes{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Function:               controllers.Login,
	AuthenticationRequires: false,
}
