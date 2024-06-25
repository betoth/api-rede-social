package routes

import (
	"api-rede-social/src/controller"
	"net/http"
)

var userRoutes = []Routes{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controller.CreateUser,
		AuthenticationRequires: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controller.SearchUsers,
		AuthenticationRequires: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Function:               controller.SearchUser,
		AuthenticationRequires: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Function:               controller.UpdateUser,
		AuthenticationRequires: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controller.DeleteUser,
		AuthenticationRequires: false,
	},
}
