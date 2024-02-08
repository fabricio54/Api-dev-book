package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// teremos duas funções para tratar os erros: uma que retorna o json do erro e a outra que envia o erro para ser formatada em json
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	// passando o status que vem no parâmetro para a api
	w.WriteHeader(statusCode)

	// pegando a reposta que nos parâmetros e passando o json
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// chamaremos a função json para trabalhar com os dados que passaremos
func Error(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct{
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}