package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Routes{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Handler:     controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Handler:     controllers.ListUsers,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Handler:     controllers.GetUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Handler:     controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Handler:     controllers.DeleteUser,
		RequireAuth: false,
	},
}
