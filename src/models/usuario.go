package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário na rede
type Usuario struct {
	ID       uint64  	`json:"id,omitempty"`
	Nome     string 	`json:"nome,omitempty"`
	Nick     string 	`json:"nick,omitempty"`
	Email    string 	`json:"email,omitempty"`
	Senha    string 	`json:"senha,omitempty"`
	CriadoEm time.Time 	`json:"criadoEm,omitempty"`
}

// Preparar vai reparar e validar as informaçõe do usuário
func (user *Usuario) Preparar(etapa string) error {
	if err := user.validar(etapa); err != nil {
		return err
	}

	if err := user.formatar("cadastro"); err != nil {
		return err
	}
	
	return nil
}

func (user *Usuario) validar(etapa string) error {
	if user.Nome == "" {
		return errors.New("o campo Nome é obrigatório")
	}

	if user.Nick == "" {
		return errors.New("o campo Nick é obrigatório")
	}

	if user.Email == "" {
		return errors.New("o campo Email deve ser preenchido")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("o email inserido é inválido")
	}

	if user.Senha == "" && etapa == "cadastro" {
		return errors.New("o campo Senha é obrigatório")
	}

	return nil
}

func (user *Usuario) formatar(etapa string) error {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nome)
	user.Email = strings.TrimSpace(user.Nome)

	if etapa == "cadastro" {
		senhaComHash, err := security.Hash(user.Senha)
		if err != nil {
			return err
		}

		user.Senha = string(senhaComHash)
	}

	return nil
}