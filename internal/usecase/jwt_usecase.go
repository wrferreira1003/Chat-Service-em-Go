package usecase

import "github.com/wrferreira1003/chat-service/internal/repository"

type JWTUsecase interface {
	ValidateToken(tokenString string) (bool, error)
}

// jwtUsecaseImpl é a implementação da interface JWTUsecase
type jwtUsecaseImpl struct {
	repository repository.JWTRepository
}

// NewJWTUsecase cria uma nova instância do caso de uso do JWT
func NewJWTUsecase(repository repository.JWTRepository) JWTUsecase {
	return &jwtUsecaseImpl{repository: repository}
}

// ValidateToken valida o token JWT fornecido
func (j *jwtUsecaseImpl) ValidateToken(tokenString string) (bool, error) {
	return j.repository.ValidateToken(tokenString)
}
