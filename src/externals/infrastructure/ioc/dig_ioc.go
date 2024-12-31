package ioc

import (
	"github.com/Ig0or/tyche/src/adapters/controllers"
	"github.com/Ig0or/tyche/src/adapters/ports/controllers_interface"
	"github.com/Ig0or/tyche/src/adapters/presenters"
	"github.com/Ig0or/tyche/src/adapters/repositories"
	"github.com/Ig0or/tyche/src/application/ports/presenters_interface"
	"github.com/Ig0or/tyche/src/application/ports/repositories_interface"
	"github.com/Ig0or/tyche/src/application/ports/use_cases_interface"
	"github.com/Ig0or/tyche/src/application/use_cases"
	"github.com/Ig0or/tyche/src/externals/infrastructure/database"
	"github.com/Ig0or/tyche/src/externals/infrastructure/logger"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/database_interface"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/logger_interface"
	"github.com/Ig0or/tyche/src/externals/ports/router_interface"
	"github.com/Ig0or/tyche/src/externals/routers"
	"go.uber.org/dig"
)

type Dependency struct {
	Implementation interface{}
	Interface      interface{}
	Name           string
}

type DigIoc struct {
	container *dig.Container
}

func NewDigIoc() *DigIoc {
	container := dig.New()

	digIoc := &DigIoc{container: container}

	return digIoc
}

func (ioc *DigIoc) LoadProviders() {
	dependencies := ioc.createDependencies()

	for _, dependency := range dependencies {
		err := ioc.container.Provide(dependency.Implementation, dig.As(dependency.Interface), dig.Name(dependency.Name))

		if err != nil {
			panic("Fail to build IOC container: " + err.Error())
		}
	}

}

func (ioc *DigIoc) Invoke(function any) {
	err := ioc.container.Invoke(function)

	if err != nil {
		panic("Fail to invoke main function inside IOC container: " + err.Error())
	}
}

func (ioc *DigIoc) createDependencies() []Dependency {
	var dependencies []Dependency

	dependencies = ioc.provideInfrastructureDependencies(dependencies)
	dependencies = ioc.provideRepositoryDependencies(dependencies)
	dependencies = ioc.providePresenterDependencies(dependencies)
	dependencies = ioc.provideUseCaseDependencies(dependencies)
	dependencies = ioc.provideControllerDependencies(dependencies)
	dependencies = ioc.provideRouterDependencies(dependencies)

	return dependencies
}

func (ioc *DigIoc) provideInfrastructureDependencies(dependencies []Dependency) []Dependency {
	infrastructureDependencies := []Dependency{
		{
			Implementation: logger.NewLogger,
			Interface:      new(logger_interface.LoggerInterface),
			Name:           "Logger",
		},
		{
			Implementation: database.NewPostgresDatabase,
			Interface:      new(database_interface.DatabaseInterface),
			Name:           "PostgresDatabase",
		},
	}

	for _, dependency := range infrastructureDependencies {
		dependencies = append(dependencies, dependency)
	}

	return dependencies
}

func (ioc *DigIoc) provideRepositoryDependencies(dependencies []Dependency) []Dependency {
	repositoryDependencies := []Dependency{
		{
			Implementation: repositories.NewAccountRepository,
			Interface:      new(repositories_interface.AccountRepositoryInterface),
			Name:           "AccountRepository",
		},
		{
			Implementation: repositories.NewTransactionRepository,
			Interface:      new(repositories_interface.TransactionRepositoryInterface),
			Name:           "TransactionRepository",
		},
	}

	for _, dependency := range repositoryDependencies {
		dependencies = append(dependencies, dependency)
	}

	return dependencies
}

func (ioc *DigIoc) providePresenterDependencies(dependencies []Dependency) []Dependency {
	presenterDependencies := []Dependency{
		{
			Implementation: presenters.NewCreateAccountPresenter,
			Interface:      new(presenters_interface.CreateAccountPresenterInterface),
			Name:           "CreateAccountPresenter",
		},
		{
			Implementation: presenters.NewTransactionPresenter,
			Interface:      new(presenters_interface.TransactionPresenterInterface),
			Name:           "TransactionPresenter",
		},
	}

	for _, dependency := range presenterDependencies {
		dependencies = append(dependencies, dependency)
	}

	return dependencies
}

func (ioc *DigIoc) provideUseCaseDependencies(dependencies []Dependency) []Dependency {
	useCaseDependencies := []Dependency{
		{
			Implementation: use_cases.NewCreateAccountUseCase,
			Interface:      new(use_cases_interface.CreateAccountUseCaseInterface),
			Name:           "CreateAccountUseCase",
		},
	}

	for _, dependency := range useCaseDependencies {
		dependencies = append(dependencies, dependency)
	}

	return dependencies
}

func (ioc *DigIoc) provideControllerDependencies(dependencies []Dependency) []Dependency {
	controllerDependencies := []Dependency{
		{
			Implementation: controllers.NewAccountController,
			Interface:      new(controllers_interface.AccountControllerInterface),
			Name:           "AccountController",
		},
	}

	for _, dependency := range controllerDependencies {
		dependencies = append(dependencies, dependency)
	}

	return dependencies
}

func (ioc *DigIoc) provideRouterDependencies(dependencies []Dependency) []Dependency {
	routerDependencies := []Dependency{
		{
			Implementation: routers.NewAccountRouter,
			Interface:      new(router_interface.RouterInterface),
			Name:           "AccountRouter",
		},
	}

	for _, dependency := range routerDependencies {
		dependencies = append(dependencies, dependency)
	}

	return dependencies
}
