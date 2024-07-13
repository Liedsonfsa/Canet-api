package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota {
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: 	"/usuarios/{usuarioId}/follow",
		Metodo: http.MethodPost,
		Funcao: controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: 	"/usuarios/{usuarioId}/unfollow",
		Metodo: http.MethodPost,
		Funcao: controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: 	"/usuarios/{usuarioId}/followers",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		URI: 	"/usuarios/{usuarioId}/following",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarQuemSigo,
		RequerAutenticacao: true,
	},
	{
		URI: 	"/usuarios/{usuarioId}/update-password",
		Metodo: http.MethodPost,
		Funcao: controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}