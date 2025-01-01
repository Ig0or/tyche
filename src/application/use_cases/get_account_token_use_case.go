package use_cases

import (
	"github.com/Ig0or/tyche/src/application/data_types/dtos"
	"github.com/Ig0or/tyche/src/application/data_types/requests"
	"github.com/Ig0or/tyche/src/application/ports/presenters_interface"
	"github.com/Ig0or/tyche/src/application/ports/repositories_interface"
	"github.com/Ig0or/tyche/src/domain/custom_errors"
	"github.com/Ig0or/tyche/src/domain/entities"
	"github.com/Ig0or/tyche/src/externals/ports/services_interface"
	"go.uber.org/dig"
)

type GetAccountTokenUseCase struct {
	accountPresenter         presenters_interface.AccountPresenterInterface
	getAccountTokenPresenter presenters_interface.GetAccountTokenPresenterInterface
	accountRepository        repositories_interface.AccountRepositoryInterface
	tokenService             services_interface.TokenServiceInterface
}

type GetAccountTokenUseCaseDependencies struct {
	dig.In

	AccountPresenter         presenters_interface.AccountPresenterInterface         `name:"AccountPresenter"`
	GetAccountTokenPresenter presenters_interface.GetAccountTokenPresenterInterface `name:"GetAccountTokenPresenter"`
	AccountRepository        repositories_interface.AccountRepositoryInterface      `name:"AccountRepository"`
	TokenService             services_interface.TokenServiceInterface               `name:"TokenService"`
}

func NewGetAccountTokenUseCase(dependencies GetAccountTokenUseCaseDependencies) *GetAccountTokenUseCase {
	getAccountTokenUseCase := &GetAccountTokenUseCase{
		accountPresenter:         dependencies.AccountPresenter,
		getAccountTokenPresenter: dependencies.GetAccountTokenPresenter,
		accountRepository:        dependencies.AccountRepository,
		tokenService:             dependencies.TokenService,
	}

	return getAccountTokenUseCase
}

func (useCase *GetAccountTokenUseCase) generateAccessToken(accountEntity *entities.AccountEntity) (string, *custom_errors.BaseCustomError) {
	tokenClaims, customError := useCase.getAccountTokenPresenter.FromEntityToTokenClaims(accountEntity)

	if customError != nil {
		return "", customError
	}

	accessToken, customError := useCase.tokenService.GenerateToken(tokenClaims)

	if customError != nil {
		return "", customError
	}

	return accessToken, nil
}

func (useCase *GetAccountTokenUseCase) getAccount(email string) (*entities.AccountEntity, *custom_errors.BaseCustomError) {
	accountModel, customError := useCase.accountRepository.GetAccountByEmail(email)

	if customError != nil {
		return nil, customError
	}

	accountEntity := useCase.accountPresenter.FromModelToEntity(accountModel)

	return accountEntity, nil
}

func (useCase *GetAccountTokenUseCase) GetAccountToken(request *requests.GetAccountTokenRequest) (*dtos.TokenDto, *custom_errors.BaseCustomError) {
	accountEntity, customError := useCase.getAccount(request.Email)

	if customError != nil {
		return nil, customError
	}

	customError = accountEntity.ValidatePassword(request.Password)

	if customError != nil {
		return nil, customError
	}

	accessToken, customError := useCase.generateAccessToken(accountEntity)

	if customError != nil {
		return nil, customError
	}

	tokenDto := useCase.getAccountTokenPresenter.CreateTokenDto(accessToken)

	return tokenDto, nil
}
