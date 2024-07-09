package routes

import (
	"api-rede-social/src/controllers"
	"net/http"
)

var publicationRoutes = []Routes{
	{
		URI:                    "/publication",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePublication,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/publications",
		Method:                 http.MethodGet,
		Function:               controllers.FindPublications,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/publication/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.FindPublicationByID,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/publication/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePublication,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/publication/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePublication,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/users/{id}/publications",
		Method:                 http.MethodGet,
		Function:               controllers.FindPublicationByUser,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/publication",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePublication,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/publication/{id}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePublication,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/publication/{id}/unlike",
		Method:                 http.MethodPost,
		Function:               controllers.UnlikePublication,
		AuthenticationRequires: true,
	},
}
