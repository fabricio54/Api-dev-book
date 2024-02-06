package main

import (
	"api/src/config"
	"api/src/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// chamando a função config.load para carregar as credencias e carregar a string de conexão com o banco
	config.Load()

	// importando o pacote server
	router := routers.ToGenerate()
	
	fmt.Printf(`Escutando Servidor na Porta %d`, config.Port)

	// colocando o servidor para rodaar
	log.Fatal(http.ListenAndServe(fmt.Sprintf(`:%d`, config.Port), router))
}
