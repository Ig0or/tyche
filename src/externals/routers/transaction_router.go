package routers

import (
	"github.com/Ig0or/tyche/src/adapters/ports/controllers_interface"
	"github.com/Ig0or/tyche/src/externals/ports/routers_interface"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type TransactionRouter struct {
	routerGroup *gin.RouterGroup
	controller  controllers_interface.TransactionControllerInterface
	middleware  routers_interface.MiddlewareInterface
}

type TransactionRouterDependencies struct {
	dig.In

	Controller controllers_interface.TransactionControllerInterface `name:"TransactionController"`
	Middleware routers_interface.MiddlewareInterface                `name:"Middleware"`
}

func NewTransactionRouter(dependencies TransactionRouterDependencies) *TransactionRouter {
	transactionRouter := &TransactionRouter{
		controller: dependencies.Controller,
		middleware: dependencies.Middleware,
	}

	return transactionRouter
}

func (router *TransactionRouter) RegisterRouter(engine *gin.Engine) {
	router.routerGroup = engine.Group("/transaction")
	router.registerRoutes()
}

func (router *TransactionRouter) registerRoutes() {
	router.routerGroup.POST("/:accountId", router.middleware.ResponseHandlerWithJwtAuthMiddleware(router.controller.CreateTransaction))
}
