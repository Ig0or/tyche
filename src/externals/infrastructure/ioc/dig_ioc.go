package ioc

import (
	"github.com/Ig0or/tyche/src/adapters/repositories/account"
	"github.com/Ig0or/tyche/src/application/ports/repositories/account"
	"github.com/Ig0or/tyche/src/externals/infrastructure/database"
	"github.com/Ig0or/tyche/src/externals/infrastructure/logger"
	"github.com/Ig0or/tyche/src/externals/ports/i_infrastructure/i_database"
	"github.com/Ig0or/tyche/src/externals/ports/i_infrastructure/i_logger"
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

func (ioc *DigIoc) Invoke(function func()) {
	err := ioc.container.Invoke(function)

	if err != nil {
		panic("Fail to invoke function inside IOC container: " + err.Error())
	}
}

func (ioc *DigIoc) createDependencies() []Dependency {
	var dependencies []Dependency

	dependencies = ioc.provideInfrastructureDependencies(dependencies)
	dependencies = ioc.provideRepositoryDependencies(dependencies)

	return dependencies
}

func (ioc *DigIoc) provideInfrastructureDependencies(dependencies []Dependency) []Dependency {
	infrastructureDependencies := []Dependency{
		{
			Implementation: logger.NewLogger,
			Interface:      new(i_logger.ILogger),
			Name:           "Logger",
		},
		{
			Implementation: database.NewPostgresDatabase,
			Interface:      new(i_database.IDatabase),
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
			Interface:      new(i_repositories.IAccountRepository),
			Name:           "AccountRepository",
		},
	}

	for _, dependency := range repositoryDependencies {
		dependencies = append(dependencies, dependency)
	}

	return dependencies
}
