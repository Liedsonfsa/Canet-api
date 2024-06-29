package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CriarUsuario cria um usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.Usuario
	if err = json.Unmarshal(corpoRequest, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Preparar(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	user.ID, err = repositorio.Criar(user)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// BuscarUsuario busca um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário."))
}

// BuscarUsuarios busca todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOrNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	users, err := repositorio.Buscar(nomeOrNick)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// AtualizarUsuario atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando o usuário."))
}

// DeletarUsuario deleta um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário."))
}
