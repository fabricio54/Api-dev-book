package users

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"fmt"

	//"os/user"

	"api/src/database"
	"api/src/models/usermodels"
	"api/src/repositories/userepositories"
)

// cadastrar usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		// por enquanto trataremos o erro assim. posteriormente vamos baixar um pacote para tratar erros
		log.Fatal(err)
	}

	// pegando o modelo pronto para inserção de usuário
	var u usermodels.User

	// verificando se tem erro
	if err = json.Unmarshal(requestBody, &u); err != nil {
		log.Fatal(err)
	}

	// conectando ao banco de dados
	db, err := database.Connection()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//criando repositório
	repositoriUser := userepositories.NewUserRepository(db)

	// passando um parâmetro modelo de usuários para o repositório de usuários
	userIDCreated, err := repositoriUser.CreateUser(u)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`Usuário com id: %v criado com sucesso!`, userIDCreated)))

}

// buscar usuário
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscar usuário"))
}

// buscar todos os usuários
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscar usuários"))
}

// deletando usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando usuarios"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deletando usuários"))
}

