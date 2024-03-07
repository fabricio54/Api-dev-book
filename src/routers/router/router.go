package router

import (
	//"api/src/authentication"
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// criando uma struct para a estrutura de rotas, onde teremos: URI, Método, Função (função da rota) e requer autenticação
type Router struct {
	URI                   string
	Method                string
	Function              func(w http.ResponseWriter, r *http.Request)
	requireAuthentication bool
}

// configurando todas as rotas da struct
func ConfigRouter(r *mux.Router) *mux.Router {

	routers := UsersRouters
	routers = append(routers, RouterLogin)

	for _, router := range routers {

		if router.requireAuthentication {
			r.HandleFunc(router.URI, middlewares.Logger(middlewares.Authenticate(router.Function)),).Methods(router.Method)
		} else {
			r.HandleFunc(router.URI,middlewares.Logger(router.Function)).Methods(router.Method)
		}
		
	}

	return r
}