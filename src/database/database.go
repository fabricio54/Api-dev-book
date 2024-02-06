package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// função de conexão com o banco de dados
func Connection() (*sql.DB, error) {
	db, err := sql.Open(config.DriveDataBase, config.StringConnection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
	
}