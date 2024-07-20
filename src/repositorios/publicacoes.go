package repositorios

import (
	"api/src/models"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes { 
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	ultimoID, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoID), nil
}

// BuscarPorID busca uma punica publicação por id
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	rows, err := repositorio.db.Query(`SELECT p.*, u.nick FROM publicacoes p INNER JOIN usuarios u ON u.id = p.autor_id WHERE p.id = ?`, publicacaoID)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer rows.Close()

	var publicacao models.Publicacao
	if rows.Next() {
		if err = rows.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); err != nil {
			return models.Publicacao{}, err
		}
	}

	return publicacao, nil
}