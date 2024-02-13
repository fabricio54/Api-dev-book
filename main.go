package main

import (
	"api/src/config"
	"api/src/routers"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// apagar depois a função init pois o secret será gerado apenas uma vez
/*
func init() {
	// gerando uma chave secret 
	key := make([]byte, 64)

	// usando pacote nativo do go rand e populando o slice com valores aleatórios
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	// temos que converter o slice de bytes em uma string. para isso utilizamos outro pacote: base64.
	stringBase64 := base64.StdEncoding.EncodeToString(key)

	fmt.Println(stringBase64)
}
*/
func main() {

	// chamando a função config.load para carregar as credencias e carregar a string de conexão com o banco
	config.Load()


	// importando o pacote server
	router := routers.ToGenerate()
	
	fmt.Printf(`Escutando Servidor na Porta %d`, config.Port)

	// colocando o servidor para rodaar
	log.Fatal(http.ListenAndServe(fmt.Sprintf(`:%d`, config.Port), router))
}
