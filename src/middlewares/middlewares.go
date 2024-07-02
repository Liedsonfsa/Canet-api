package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"fmt"
	"log"
	"net/http"
)

// Logger mostra as informações da rota
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Autenticar verifica se o usuário está autenticado
func Autenticar(next http.HandlerFunc)  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.TokenValidation(r); err != nil {
			responses.Erro(w, http.StatusUnauthorized, err)
			return
		}
		fmt.Println("Validando...")
		next(w, r)
	}

}