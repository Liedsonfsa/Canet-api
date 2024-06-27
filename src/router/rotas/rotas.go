package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da api
type Rota struct {
	URI    				string
	Metodo 				string
	Funcao 				func(http.ResponseWriter, *http.Request)
	RequerAutenticacao 	bool
}

// Configurar coloca todos as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
} 
