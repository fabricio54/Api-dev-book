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

// Método para retornar usuários pelo filtro do nome ou nick
func (repositori Usuarios) GetAllUsers(nameOrnick string) (*usermodels.User, error) {
	// executando a query e pegando os resultados
	resultsRows, err := repositori.db.Query("SELECT name, nick, email FROM users WHERE name=? OR nick=?", nameOrnick, nameOrnick)

	if err != nil {
		return nil, err
	}

	var user usermodels.User

	if err := resultsRows.Scan(&user.Name, &user.Nick, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
