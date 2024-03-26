package userepositories

import (
	"api/src/models/usermodels"
	"database/sql"

	"fmt"
	"time"
)

// criando uma estrutura para receber uma conexão
type Usuarios struct {
	db *sql.DB
}

// criando uma função para criar uma instância de usuários e retornar. esse processo é feito para que possamos manipular as funções específicas para usuários
func NewUserRepository(db *sql.DB) *Usuarios {
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
func (repositori Usuarios) GetAllUsers(nameOrnick string) ([]usermodels.User, error) {

	// queremos pegar não só o valor literal mais sim se alguma parte que veio na string
	nameOrnick = fmt.Sprintf("%%%s%%", nameOrnick)

	// executando a query e pegando os resultados
	rows, err := repositori.db.Query("SELECT id, name, nick, email, createdIn FROM users WHERE name OR nick LIKE ?", nameOrnick)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user []usermodels.User
	var id, name, nick, email string
	var createdIn time.Time

	for rows.Next() {
		if err = rows.Scan(&id, &name, &nick, &email, &createdIn); err != nil {
			return nil, err
		}
		user = append(user, usermodels.User{Name: name, Nick: nick, Email: email, ID: id, CreatedIn: createdIn})
	}
	/*
		if err := resultsRows.Scan(&user.Name, &user.Nick, &user.Email); err != nil {
			return nil, err
		}
	*/

	return user, nil
}

// Método retornar um usuário específico pelo id
func (repositori Usuarios) GetUser(id int) (usermodels.User, error) {
	// criando variável para pegar struct
	var user usermodels.User

	// executando e pegando os resultados
	if err := repositori.db.QueryRow("SELECT id, name, nick, email FROM users WHERE id=?", id).Scan(&user.ID, &user.Name, &user.Nick, &user.Email); err != nil {
		return usermodels.User{}, err
	}

	// retornando os dados que vieram no resultado da query
	return user, nil
}

// Método para atualizar informações de usuário no banco
func (repositori Usuarios) UpdateUser(id uint64, user usermodels.User) error {
	// criando o statement
	statement, err := repositori.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")

	if err != nil {
		return err
	}
	defer statement.Close()

	// executando statement
	if _, err = statement.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}

	return nil
}

// Método para deletar informações de usuário no banco
func (repositori Usuarios) DeleteUser(id uint64) error {
	// criando o statement
	statement, err := repositori.db.Prepare("DELETE FROM users WHERE id = ?")

	if err != nil {
		return err
	}

	// executando statement
	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

// Método para buscar usuário no banco por email
func (repositori Usuarios) SearchEmailUser(email string) (usermodels.User, error) {

	//criando modelo de user
	var user usermodels.User

	// buscar usuário no banco
	if err := repositori.db.QueryRow("SELECT id, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Password); err != nil {
		return usermodels.User{}, nil
	}

	// retornando estrutura
	return user, nil
}

// Método para seguir outro usuário
func (repositori Usuarios) Follower(idUser, idFollower uint64) error {
	statement, err := repositori.db.Prepare("insert ignore into followers (user_id, follower_id) values (?, ?)")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(idUser, idFollower); err != nil {
		return err
	}

	return nil
}

// Método para deixar de seguir usuário
func (repositori Usuarios) Unfollow(idFollowed, idFollower uint64) error {
	statement, err := repositori.db.Prepare("delete from followers where user_id = ? && follower_id = ?")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(idFollowed, idFollower); err != nil {
		return err
	}

	return nil
}

// Método para buscar seguidores
func (repositori Usuarios) SearchForFollowers(idUser uint64) ([]usermodels.User, error) {
	rows, err := repositori.db.Query(`
	SELECT u.id, u.name, u.nick, u.email, u.createdIn FROM users as u inner join followers as f
    ON u.id = f.user_id
    WHERE f.follower_id = ?
	`, idUser)

	if err != nil {
		return nil, err
	}

	// pegar todos os usuários e passar para um struct de usuários
	var users []usermodels.User
	for rows.Next() {

		var id, name, nick, email string
		var createdIn time.Time

		if err := rows.Scan(&id, &name, &nick, &email, &createdIn); err != nil {
			return nil, err
		}
		users = append(users, usermodels.User{ID: id, Name: name, Nick: nick, Email: email, CreatedIn: createdIn})
	}

	return users, nil
}

// Método para buscar usuários que te seguem
func (repositori Usuarios) GetForFollowing(idUser uint64) ([]usermodels.User, error) {
	// criando a query
	rows, err := repositori.db.Query(`
	SELECT 
		u.id, u.name, u.nick, u.email, u.createdIn 
	FROM 
		users as u inner join followers as f
    ON 
		u.id = f.follower_id
    WHERE 
		f.user_id = ?;
	`, idUser)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []usermodels.User

	for rows.Next() {

		var user usermodels.User

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}

// busca senha do usário
func (repositori Usuarios) SearchPassword(idUser uint64) (string, error) {

	var password string
	
	if err := repositori.db.QueryRow(`
		select password from users where id = ?
	`, idUser).Scan(&password); err != nil {
		return "", nil
	}

	return password, nil
}

// atualiza senha de usuário
func (repositori Usuarios) UpdatePasswordUser(idUser uint64, password string) error {
	statement, err := repositori.db.Prepare(`
		UPDATE users SET password = ? where id = ?
	`)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, idUser); err != nil {
		return err
	}

	return nil
}
