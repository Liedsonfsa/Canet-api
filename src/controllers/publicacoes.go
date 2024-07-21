package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacao models.Publicacao
	if err = json.Unmarshal(corpoRequisicao, &publicacao); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	publicacao.AutorID = usuarioId

	if err = publicacao.Preparar(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, err = repositorio.Criar(publicacao)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publicacao)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, err := repositorio.Buscar(usuarioID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publicacoes)
}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(parameters["publicacaoId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publicacao)
}

func AtualizarPublicao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(parameters["publicacaoId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoNoBanco, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacaoNoBanco.AutorID != usuarioID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar uma publicação que não seja sua"))
		return
	}

	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacao models.Publicacao
	if err = json.Unmarshal(corpoRequisicao, &publicacao); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = publicacao.Preparar(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repositorio.Atualizar(publicacaoID, publicacao); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}