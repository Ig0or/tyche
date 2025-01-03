package http_server

import (
	"github.com/Ig0or/tyche/src/externals/ports/routers_interface"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"os"
)

type GinHttpServer struct {
	engine *gin.Engine

	accountRouter     routers_interface.RouterInterface
	transactionRouter routers_interface.RouterInterface
}

type GinHttpServerDependencies struct {
	dig.In

	AccountRouter     routers_interface.RouterInterface `name:"AccountRouter"`
	TransactionRouter routers_interface.RouterInterface `name:"TransactionRouter"`
}

func NewGinHttpServer(dependencies GinHttpServerDependencies) {
	engine := gin.Default()

	httpServer := &GinHttpServer{
		engine:            engine,
		accountRouter:     dependencies.AccountRouter,
		transactionRouter: dependencies.TransactionRouter,
	}

	httpServer.registerRouters()
	httpServer.startServer()
}

func (server *GinHttpServer) registerRouters() {
	server.accountRouter.RegisterRouter(server.engine)
	server.transactionRouter.RegisterRouter(server.engine)
}

func (server *GinHttpServer) startServer() {
	serverPort := os.Getenv("SERVER_PORT")

	server.engine.Run(serverPort)
}
