package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositorios"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
	"fmt"
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

	fmt.Printf("Senha com hash: %s\n", userOnDB.Senha)

	if err = security.VerificarSenha(userOnDB.Senha, user.Senha); err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CriarToken(userOnDB.ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	
	w.Write([]byte(token))
}