package userepositories

import (
	"api/src/models/usermodels"
	"database/sql"
)

// criando uma estrutura para receber uma conexão
type Usuarios struct {
	db *sql.DB
}

// criando uma função para criar uma instância de usuários e retornar. esse processo é feito para que possamos manipular as funções específicas para usuários
func NewUserRepository(db *sql.DB) (*Usuarios) {
	return &Usuarios{db}
}

// Método para criar usuário
func (u Usuarios) CreateUser(user usermodels.User) (uint64, error) {
	return 0, nil
}

