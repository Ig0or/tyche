package services

import (
	"fmt"
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

	accessToken, err := token.SignedString(secretKeyBytes)

	if err != nil {
		customError := custom_errors.NewInternalServerError("Error while trying to generate token in JwtTokenService.", err)

		return "", customError
	}
	return accessToken, nil
}

func (service *JwtTokenService) parseToken(token *jwt.Token) (interface{}, error) {
	secretKeyBytes := []byte(service.secretKey)

	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		return nil, err
	}

	return secretKeyBytes, nil
}

func (service *JwtTokenService) ValidateToken(accessToken string) (jwt.MapClaims, *custom_errors.BaseCustomError) {
	token, err := jwt.Parse(accessToken, service.parseToken)

	if err != nil {
		customError := custom_errors.NewUnauthorizedError("Fail to validate JWT because this token is invalid or expired.", err)

		return nil, customError
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}
