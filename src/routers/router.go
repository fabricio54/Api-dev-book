package routers

import (
	"api/src/routers/router"

	"github.com/gorilla/mux"
)

// função criada para gerar rotas: criamos uma instância da função mux.router e depois passamos para a função do pacote router para configurar elas e ja retornar
func toGenerate() (*mux.Router) {
	r :=  mux.NewRouter()
	return router.ConfigRouter(r)
}


