package main

import (
	"github.com/Ig0or/tyche/src/externals/infrastructure/http_server_config"
	"github.com/Ig0or/tyche/src/externals/infrastructure/ioc_config"
	"github.com/Ig0or/tyche/src/externals/ports/i_infrastructure/i_ioc_config"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic("Fail to initialize while trying to load env: " + err.Error())
	}
}

func loadIoc() i_ioc_config.IIocConfig {
	iocContainer := ioc_config.NewDigIocConfig()
	iocContainer.LoadProviders()

	return iocContainer
}

func startApp(iocContainer i_ioc_config.IIocConfig) {
	iocContainer.Invoke(http_server_config.StartHttpServer)

}

func main() {
	loadEnv()
	iocContainer := loadIoc()
	startApp(iocContainer)
}
