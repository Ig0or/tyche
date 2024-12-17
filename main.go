package main

import (
	"github.com/Ig0or/tyche/src/externals/infrastructure/http_server"
	"github.com/Ig0or/tyche/src/externals/infrastructure/ioc"
	"github.com/Ig0or/tyche/src/externals/ports/i_infrastructure/i_ioc"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic("Fail to initialize while trying to load env: " + err.Error())
	}
}

func loadIoc() i_ioc.IIoc {
	iocContainer := ioc.NewDigIoc()
	iocContainer.LoadProviders()

	return iocContainer
}

func startApp(iocContainer i_ioc.IIoc) {
	iocContainer.Invoke(http_server.StartHttpServer)

}

func main() {
	loadEnv()
	iocContainer := loadIoc()
	startApp(iocContainer)
}
