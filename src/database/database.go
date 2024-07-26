package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Conectar cria, estabelece e retorna uma conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConexao)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}