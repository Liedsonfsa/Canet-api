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

// Buscar buscar todas as publicações do usuário e das pessoas que ele segue
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]models.Publicacao, error) {
	rows, err := repositorio.db.Query(`select distinct p.*, u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id inner join seguidores s on p.autor_id = s.usuario_id where u.id = ? or seguidor_id = ? order by 1 desc`, usuarioID, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publicacoes []models.Publicacao

	for rows.Next() {
		var publicacao models.Publicacao

		if err = rows.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, err := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}

	return nil
}
