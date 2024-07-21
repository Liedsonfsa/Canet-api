package models

// Senha armazena as senhas nova e atual na hora da mudança de senha
type Senha struct {
	Nova string `json:"nova"`
	Atual string `json:"atual"`
}
