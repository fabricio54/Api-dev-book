package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando Api")

	// importando o pacote server
	router := router.Gerar()

	// colocando o servidor para rodaar
	log.Fatal(http.ListenAndServe(":5000", router))
}
