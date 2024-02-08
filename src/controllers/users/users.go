package users

import (
	"encoding/json"
	//"fmt"
	"io"
	//"log"
	"net/http"

	//"os/user"

	"api/src/database"
	"api/src/models/usermodels"
	"api/src/repositories/userepositories"
	"api/src/response"
)

// cadastrar usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		// por enquanto trataremos o erro assim. posteriormente vamos baixar um pacote para tratar erros
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// pegando o modelo pronto para inserção de usuário
	var u usermodels.User

	// verificando se tem erro
	if err = json.Unmarshal(requestBody, &u); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	// conectando ao banco de dados
	db, err := database.Connection()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	//criando repositório
	repositoriUser := userepositories.NewUserRepository(db)

	// passando um parâmetro modelo de usuários para o repositório de usuários
	userIDCreated, err := repositoriUser.CreateUser(u)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	// como criamos uma estrutura para tratar os erros agora chamamos uma para enviar as repostas em json
	response.JSON(w, http.StatusOK, struct {
		ID uint64 `json:"id"`
	}{
		ID: userIDCreated,
	})

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

