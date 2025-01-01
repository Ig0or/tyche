package presenters

import (
	"github.com/Ig0or/tyche/src/application/data_types/dtos"
	"github.com/Ig0or/tyche/src/application/data_types/requests"
	"github.com/Ig0or/tyche/src/application/data_types/responses"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strconv"
	"time"
)

type GetAccountTokenPresenter struct {
}

func NewGetAccountTokenPresenter() *GetAccountTokenPresenter {
	return &GetAccountTokenPresenter{}
}

func (presenter *GetAccountTokenPresenter) FromContextToRequest(context *gin.Context) (*requests.GetAccountTokenRequest, *custom_errors.BaseCustomError) {
	var request requests.GetAccountTokenRequest

	err := context.ShouldBindJSON(&request)

	if err != nil {
		message := custom_errors.FormatRequestValidationError(err)
		customError := custom_errors.NewBadRequestError(message, err)

		return nil, customError
	}

	return &request, nil
}

func (presenter *GetAccountTokenPresenter) FromEntityToTokenClaims(entity *entities.AccountEntity) (jwt.MapClaims, *custom_errors.BaseCustomError) {
	hoursToLive, err := strconv.Atoi(os.Getenv("JWT_TTL_IN_HOURS"))

	if err != nil {
		customError := custom_errors.NewInternalServerError("Error while trying to create jwt claims in GetAccountTokenPresenter.", err)

		return nil, customError
	}

	timeToLive := time.Now().UTC().Add(time.Duration(hoursToLive) * time.Hour)

	expirationTime := jwt.NewNumericDate(timeToLive)

	claims := jwt.MapClaims{
		"account_id": entity.AccountId,
		"email":      entity.Email,
		"exp":        expirationTime,
	}

	return claims, nil
}

func (presenter *GetAccountTokenPresenter) CreateTokenDto(accessToken string) *dtos.TokenDto {
	tokenDto := &dtos.TokenDto{AccessToken: accessToken}

	return tokenDto
}

func (presenter *GetAccountTokenPresenter) FromDtoToResponse(dto *dtos.TokenDto) *responses.BaseApiResponse {
	accountTokenResponse := &responses.GetAccountTokenResponse{
		AccessToken: dto.AccessToken,
	}

	apiResponse := &responses.BaseApiResponse{
		Payload:    accountTokenResponse,
		Success:    true,
		StatusCode: http.StatusCreated,
	}

	return apiResponse
}
