package main

import (
	"github.com/Ig0or/tyche/src/externals/infrastructure/http_server"
	"github.com/Ig0or/tyche/src/externals/infrastructure/ioc"
	"github.com/Ig0or/tyche/src/externals/ports/infrastructure/ioc_interface"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic("Fail to initialize while trying to load env: " + err.Error())
	}
}

func loadIoc() ioc_interface.IocInterface {
	iocContainer := ioc.NewDigIoc()
	iocContainer.LoadProviders()

	return iocContainer
}

func startApp(iocContainer ioc_interface.IocInterface) {
	iocContainer.Invoke(http_server.NewGinHttpServer)

}

func main() {
	loadEnv()
	iocContainer := loadIoc()
	startApp(iocContainer)
}
