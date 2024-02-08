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
func (repositori Usuarios) CreateUser(user usermodels.User) (uint64, error) {
	// preparando o statement
	statement, err := repositori.db.Prepare(`INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	// executando o statement
	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastIDEntered, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastIDEntered), nil
}
