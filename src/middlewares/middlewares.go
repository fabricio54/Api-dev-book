package middlewares

import (
	"api/src/authentication"
	"api/src/response"
	"log"
	"net/http"
)

// middlewares são criados para ficar entre as camadas requisição e reposta

// função para mostrar informações das rotas que forma acessadas
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// função que verifica se usuário está autenticado
func Authenticate(next http.HandlerFunc) (http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidatToken(r); err != nil {
			response.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}