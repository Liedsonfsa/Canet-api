package rotas

import (
	"net/http"
	"api/src/controllers"
)

var rotasPublicacoes = []Rota {
	{
		URI: "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{publicacaoId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{publicacaoId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarPublicao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{publicacaoId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
}