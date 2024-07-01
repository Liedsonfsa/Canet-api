package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken cria um token com as permições para o usuário
func CriarToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["usuarioId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("Secret"))
}