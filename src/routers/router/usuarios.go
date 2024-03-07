package router

import (
	"api/src/controllers/users"
	"net/http"
)

// construindo as 5 principais rotas nessa variável
var UsersRouters = []Router {
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: users.CreateUser,
		requireAuthentication: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Function: users.GetAllUsers,
		requireAuthentication: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Function: users.GetUser,
		requireAuthentication: false,
	}, 
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Function: users.UpdateUser,
		requireAuthentication: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Function: users.DeleteUser,
		requireAuthentication: false,
	},
}