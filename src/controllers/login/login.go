package login

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models/usermodels"
	"api/src/repositories/userepositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// Função para logar usuário na api
func Login(w http.ResponseWriter, r *http.Request) {
	// pegando o corpo da requisição
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// criando estrutura para usuário
	var user usermodels.User

	// passando o json para a estrutura
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	// abrindo conexão com o banco
	db, err := database.Connection()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// criando repositorio para login
	repositori := userepositories.NewUserRepository(db)

	// função para buscar usuário na base de dados
	userInternalBd, err := repositori.SearchEmailUser(user.Email)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	// comparando a senha que veio no userInternalBd com a senha que veio na requisição
	if err = security.CheckPassword(user.Password, userInternalBd.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	idUser, err := strconv.ParseUint(userInternalBd.ID, 10, 64)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// chamando a função para gerar token
	token, _ := authentication.CreateTokenJWT(idUser)

	w.Write([]byte(token))
}
