package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = user.Preparar("cadastro"); err != nil {
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
	parameters := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)

	if err != nil {
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
	user, err := repositorio.BuscarPorID(usuarioID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
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
	parameters := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	fmt.Println(tokenUserID)
	if tokenUserID != usuarioID {
		responses.Erro(w, http.StatusForbidden, errors.New("você não tem autorização para autorizar esse usuário"))
		return
	}

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

	if err = user.Preparar("edição"); err != nil {
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
	if err = repositorio.Atualizar(usuarioID, user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletarUsuario deleta um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != tokenUserID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possível deletar esse usuário"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if err = repositorio.Deletar(usuarioID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// SeguirUsuario estabeleca uma conexão entre dois usuário
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if seguidorID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if err = repositorio.Seguir(userID, seguidorID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// PararDeSeguirUsuario para de seguir um usuário
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if seguidorID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possível parar de seguir você mesmo"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if err = repositorio.PararDeSeguir(userID, seguidorID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeguidores
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if err != nil {
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
	seguidores, err := repositorio.BuscarSeguidores(usuarioID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, seguidores)
}