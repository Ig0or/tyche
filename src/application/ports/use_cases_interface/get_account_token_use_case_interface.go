package use_cases_interface

import (
	"github.com/Ig0or/tyche/src/application/data_types/dtos"
	"github.com/Ig0or/tyche/src/application/data_types/requests"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
)

type GetAccountTokenUseCaseInterface interface {
	GetAccountToken(request *requests.GetAccountTokenRequest) (*dtos.TokenDto, *custom_errors.BaseCustomError)
}
