package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

// Login é responsável por autenticar o usuário
func Login(w http.ResponseWriter, r *http.Request) {
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
	
	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	userOnDB, err := repositorio.BuscarPorEmail(user.Email)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = security.VerificarSenha(user.Senha, userOnDB.Senha); err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	w.Write([]byte("Você está logado"))
}