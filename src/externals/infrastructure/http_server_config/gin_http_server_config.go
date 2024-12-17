package http_server_config

import (
	"github.com/Ig0or/tyche/src/externals/ports/i_router"
	"github.com/Ig0or/tyche/src/externals/router"
	"github.com/gin-gonic/gin"
	"os"
)

func registerRouters(engine *gin.Engine) {
	var allRouters []i_router.IRouter

	accountRouter := router.NewAccountRouter(engine)

	allRouters = append(allRouters, accountRouter)

	for _, router := range allRouters {
		router.RegisterRoutes()
	}

}

func StartHttpServer() {
	serverPort := os.Getenv("SERVER_PORT")

	engine := gin.Default()

	registerRouters(engine)

	engine.Run(serverPort)
}
