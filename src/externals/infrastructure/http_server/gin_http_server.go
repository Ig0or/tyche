package http_server

import (
	"github.com/Ig0or/tyche/src/externals/ports/router_interface"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"os"
)

type GinHttpServer struct {
	engine *gin.Engine

	accountRouter router_interface.RouterInterface
}

type GinHttpServerDependencies struct {
	dig.In

	AccountRouter router_interface.RouterInterface `name:"AccountRouter"`
}

func NewGinHttpServer(dependencies GinHttpServerDependencies) {
	engine := gin.Default()

	httpServer := &GinHttpServer{engine: engine, accountRouter: dependencies.AccountRouter}

	httpServer.registerRouters()
	httpServer.startServer()
}

func (server *GinHttpServer) registerRouters() {
	server.accountRouter.RegisterRouter(server.engine)
}

func (server *GinHttpServer) startServer() {
	serverPort := os.Getenv("SERVER_PORT")

	server.engine.Run(serverPort)
}
