package router

import (
	"api/src/controllers/login"
	"net/http"
)

var RouterLogin = Router{
	URI:    "/login",
	Method: http.MethodPost,
	Function: login.Login,
	requireAuthentication: false,
}
