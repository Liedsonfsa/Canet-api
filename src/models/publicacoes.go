package models

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicação feita por um usuário
type Publicacao struct {
	ID        uint64 	`json:"id,omitempty"`
	Titulo    string 	`json:"titulo,omitempty"`
	Conteudo  string 	`json:"conteudo,omitempty"`
	AutorID   uint64 	`json:"autorId,omitempty"`
	AutorNick string 	`json:"autorNick,omitempty"`
	Curtidas  uint64 	`json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

// Preparar cai validar e formatar a publicação recebida
func (publicacao *Publicacao) Preparar() error {
	if err := publicacao.validar(); err != nil {
		return err
	}

	publicacao.formatar()
	return nil
}

// validar verifica se existe algum conteúdo na publicação
func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("por favor coloque um título na sua publicação")
	}

	if publicacao.Conteudo == "" {
		return errors.New("a publicação presica possuir algum tipo de conteúdo")
	}

	return nil
}

// formatar formata a publicação
func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
