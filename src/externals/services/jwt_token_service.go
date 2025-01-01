package services

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type JwtTokenService struct {
	secretKey string
}

func NewJwtTokenService() *JwtTokenService {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	jwtTokenService := &JwtTokenService{secretKey: secretKey}

	return jwtTokenService
}

func (service *JwtTokenService) GenerateToken(claims jwt.MapClaims) (string, *custom_errors.BaseCustomError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKeyBytes := []byte(service.secretKey)

	accessTokenString, err := token.SignedString(secretKeyBytes)

	if err != nil {
		customError := custom_errors.NewInternalServerError("Error while trying to generate token in JwtTokenService", err)

		return "", customError
	}
	return accessTokenString, nil
}
