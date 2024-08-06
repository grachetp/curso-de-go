package routes

import (
	"api/src/router/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.FindUsers,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.FindUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
