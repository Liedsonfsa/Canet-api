package repositorios

import (
	"database/sql"
	"api/src/models"
)

type usuarios struct {
	db *sql.DB
}

// NovoepositorioDeUsuarios cria um novo repositório de usuario
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// CriarUsuario cria um usuário no banco de dados
func (repositorio usuarios) Criar(user models.Usuario) (uint64, error){
	statement, err := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES(?, ?, ?, ?)")

	if err != nil {
		return 0, nil
	}

	defer statement.Close()
	result, err := statement.Exec(user.Nome, user.Nick, user.Email, user.Senha)

	if err != nil {
		return 0, nil
	}

	ultimoID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(ultimoID), nil
}
