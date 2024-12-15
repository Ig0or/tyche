package ioc_config

import (
	"github.com/Ig0or/tyche/src/externals/infrastructure/database_config"
	"github.com/Ig0or/tyche/src/externals/infrastructure/logger_config"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/i_database_config"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/i_logger_config"
	"github.com/Ig0or/tyche/src/server"
	"go.uber.org/dig"
)

type Dependency struct {
	Constructor interface{}
	Interface   interface{}
	Name        string
}

type DigIocConfig struct {
}

func NewDigIocConfig() *DigIocConfig {
	return &DigIocConfig{}
}

func (iocConfig *DigIocConfig) BuildIOCContainer() {
	container := dig.New()

	iocConfig.loadProviders(container)

	err := container.Invoke(server.NewLero)

	if err != nil {
		panic("Fail to invoke server using IOC container: " + err.Error())
	}

}

func (iocConfig *DigIocConfig) loadProviders(container *dig.Container) {
	dependencies := iocConfig.createDependencies()

	for _, dependency := range dependencies {
		err := container.Provide(dependency.Constructor, dig.As(dependency.Interface), dig.Name(dependency.Name))

		if err != nil {
			panic("Fail to build IOC container: " + err.Error())
		}
	}

}

func (iocConfig *DigIocConfig) createDependencies() []Dependency {
	var dependencies []Dependency

	dependencies = iocConfig.provideInfrastructureDependencies(dependencies)

	return dependencies
}

func (iocConfig *DigIocConfig) provideInfrastructureDependencies(dependencies []Dependency) []Dependency {
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
