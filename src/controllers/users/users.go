package users

import (
	"encoding/json"
	"strconv"
	"strings"

	//"fmt"
	"io"
	//"log"
	"net/http"

	//"os/user"

	"github.com/gorilla/mux"
	"api/src/database"
	"api/src/models/usermodels"
	"api/src/repositories/userepositories"
	"api/src/response"
)

// observação: trabalhar com queryParams: podemos trabalhar na própria url r ja para pegar dados da propria url usamos o pacote mux 

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

	// validações nos dados recebidos
	if err = u.Prepare(); err != nil {
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
	// pegando o id direto da url
	id := mux.Vars(r)["id"]
	
	// pegando parâmetro id
	idUser, err := strconv.Atoi(id)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	// abrindo uma nova conexão
	db, err := database.Connection()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// criando novo repositório
	repositoryUser := userepositories.NewUserRepository(db)

	// utilizando o repositório para buscar usuário
	user, err := repositoryUser.GetUser(idUser)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	// passando o resultado em json
	response.JSON(w, http.StatusOK, user)
}

// buscar todos os usuários por filtro de nome ou nick
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// primeiro pegamos o parâmetro usuario vindo da rota:
	NameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	// conexão com o banco de dados
	db, err := database.Connection()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// criando um novo repositori para usuário
	repositoriUser := userepositories.NewUserRepository(db)

	// pegando os resultados para nome ou nick passados
	users, err := repositoriUser.GetAllUsers(NameOrNick)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}

	// passando os resultados de usuários com nomes ou nicks correspondentes
	response.JSON(w, http.StatusOK, users)
}

// deletando usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando usuarios"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deletando usuários"))
}

