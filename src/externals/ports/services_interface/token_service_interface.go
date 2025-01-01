package services_interface

import (
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/golang-jwt/jwt/v5"
)

type TokenServiceInterface interface {
	GenerateToken(claims jwt.MapClaims) (string, *custom_errors.BaseCustomError)
}
