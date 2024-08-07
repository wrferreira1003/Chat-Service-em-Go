package repository

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/wrferreira1003/chat-service/internal/config"
)

type JWTRepository interface {
	ValidateToken(tokenString string) (bool, error) // Valida o token
}

// Implementação da interface JWTRepository
type jwtRepositoryIml struct {
	secretKey []byte
}

func NewJWTRepository() JWTRepository {
	return &jwtRepositoryIml{
		secretKey: []byte(config.GetToken()), //TODO: verificar se a chave secreta está correta na variável de ambiente
	}
}

// Validar o token conforme a chave secreta
func (j *jwtRepositoryIml) ValidateToken(tokenString string) (bool, error) {
	log.Printf("Validating token: %s", tokenString)
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) // token.Header["alg"] é o método de assinatura do token
		}
		return j.secretKey, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return false, err
	}

	return token.Valid, nil
}
