package security

import "golang.org/x/crypto/bcrypt"

// Hash faz o hash de uma string
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha),bcrypt.DefaultCost)
}

// VerificarSenha compara se uma senha e um hash são iguais
func VerificarSenha(senhaHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}