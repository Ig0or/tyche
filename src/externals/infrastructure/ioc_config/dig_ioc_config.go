package ioc_config

import (
	"github.com/Ig0or/tyche/src/adapters/repositories/account"
	"github.com/Ig0or/tyche/src/application/ports/repositories/account"
	"github.com/Ig0or/tyche/src/externals/infrastructure/database_config"
	"github.com/Ig0or/tyche/src/externals/infrastructure/logger_config"
	"github.com/Ig0or/tyche/src/externals/ports/i_infrastructure/i_database_config"
	"github.com/Ig0or/tyche/src/externals/ports/i_infrastructure/i_logger_config"
	"go.uber.org/dig"
)

type Dependency struct {
	Constructor interface{}
	Interface   interface{}
	Name        string
}

type DigIocConfig struct {
	container *dig.Container
}

func NewDigIocConfig() *DigIocConfig {
	container := dig.New()

	digIoc := &DigIocConfig{container: container}

	return digIoc
}

func (ioc *DigIocConfig) LoadProviders() {
	dependencies := ioc.createDependencies()

	for _, dependency := range dependencies {
		err := ioc.container.Provide(dependency.Constructor, dig.As(dependency.Interface), dig.Name(dependency.Name))

		if err != nil {
			panic("Fail to build IOC container: " + err.Error())
		}
	}

}

func (ioc *DigIocConfig) Invoke(function func()) {
	err := ioc.container.Invoke(function)

	if err != nil {
		panic("Fail to invoke function inside IOC container: " + err.Error())
	}
}

func (ioc *DigIocConfig) createDependencies() []Dependency {
	var dependencies []Dependency

	dependencies = ioc.provideInfrastructureDependencies(dependencies)
	dependencies = ioc.provideRepositoryDependencies(dependencies)

	return dependencies
}

func (ioc *DigIocConfig) provideInfrastructureDependencies(dependencies []Dependency) []Dependency {
	infrastructureDependencies := []Dependency{
		{
			Constructor: logger_config.NewLoggerConfig,
			Interface:   new(i_logger_config.ILoggerConfig),
			Name:        "LoggerConfig",
		},
		{
			Constructor: database_config.NewPostgresDatabaseConfig,
			Interface:   new(i_database_config.IDatabaseConfig),
			Name:        "PostgresDatabaseConfig",
		},
	}

	for _, dependency := range infrastructureDependencies {
		dependencies = append(dependencies, dependency)
	}

	return dependencies
}

func (ioc *DigIocConfig) provideRepositoryDependencies(dependencies []Dependency) []Dependency {
	repositoryDependencies := []Dependency{
		{
			Constructor: repositories.NewAccountRepository,
			Interface:   new(i_repositories.IAccountRepository),
			Name:        "AccountRepository",
		},
	}

	for _, dependency := range repositoryDependencies {
		dependencies = append(dependencies, dependency)
	}

	return dependencies
}
