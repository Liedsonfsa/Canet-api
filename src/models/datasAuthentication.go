package models

// DatasAuthentication contém o token e o id do usuário autenticado
type DatasAuthentication struct {
	ID 		string `json:"id"`
	Token 	string `json:"token"`
}