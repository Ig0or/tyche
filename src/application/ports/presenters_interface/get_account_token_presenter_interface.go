package presenters_interface

import (
	"github.com/Ig0or/tyche/src/application/data_types/dtos"
	"github.com/Ig0or/tyche/src/application/data_types/requests"
	"github.com/Ig0or/tyche/src/application/data_types/responses"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type GetAccountTokenPresenterInterface interface {
	FromContextToRequest(context *gin.Context) (*requests.GetAccountTokenRequest, *custom_errors.BaseCustomError)
	FromEntityToTokenClaims(entity *entities.AccountEntity) (jwt.MapClaims, *custom_errors.BaseCustomError)
	CreateTokenDto(accessToken string) *dtos.TokenDto
	FromDtoToResponse(dto *dtos.TokenDto) *responses.BaseApiResponse
}
