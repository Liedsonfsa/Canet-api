package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// NovoepositorioDeUsuarios cria um novo repositório de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere um usuário no banco de dados
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

// Buscar busca os usuários com um determinado nick ou nome
func (repositorio usuarios) Buscar(nomeOrNick string) ([]models.Usuario, error) {
	nomeOrNick = fmt.Sprintf("%%%s%%", nomeOrNick)

	rows, err := repositorio.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? or nick LIKE ?", nomeOrNick, nomeOrNick)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.Usuario

	for rows.Next() {
		var user models.Usuario
		if err = rows.Scan(&user.ID, &user.Nome, &user.Nick, &user.Email, &user.CriadoEm); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// BuscarPorID busca um único usuário por ID
func (repositorio *usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	rows, err := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", ID,
	)

	if err != nil {
		return models.Usuario{}, err
	}
	defer rows.Close()
	
	var user models.Usuario
	if rows.Next() {
		if err = rows.Scan(&user.ID, &user.Nome, &user.Nick, &user.Email, &user.CriadoEm); err != nil {
			return models.Usuario{}, err
		}
	}

	return user, nil
}

// Atualizar atualiza um usuário no banco de dados
func (repositorio *usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, err := repositorio.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}

	return nil
}

// Deletar deleta as informações de um usuário no banco de dados
func (repositorio *usuarios) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}