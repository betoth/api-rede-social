package routes

import (
	"api-rede-social/src/controllers"
	"net/http"
)

var userRoutes = []Routes{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		AuthenticationRequires: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUsers,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUser,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/users/{id}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/users/{id}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnfollowUser,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/users/{id}/followers",
		Method:                 http.MethodGet,
		Function:               controllers.SearchFollowers,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/users/{id}/following",
		Method:                 http.MethodGet,
		Function:               controllers.SearchFollowing,
		AuthenticationRequires: true,
	},
	{
		URI:                    "/users/{id}/update-password",
		Method:                 http.MethodPost,
		Function:               controllers.UpdatePassword,
		AuthenticationRequires: true,
	},
}
