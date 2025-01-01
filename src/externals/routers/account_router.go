package routers

import (
	"github.com/Ig0or/tyche/src/adapters/ports/controllers_interface"
	"github.com/Ig0or/tyche/src/externals/ports/routers_interface"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type AccountRouter struct {
	routerGroup *gin.RouterGroup
	controller  controllers_interface.AccountControllerInterface
	middleware  routers_interface.MiddlewareInterface
}

type AccountRouterDependencies struct {
	dig.In

	Controller controllers_interface.AccountControllerInterface `name:"AccountController"`
	Middleware routers_interface.MiddlewareInterface            `name:"Middleware"`
}

func NewAccountRouter(dependencies AccountRouterDependencies) *AccountRouter {
	accountRouter := &AccountRouter{
		controller: dependencies.Controller,
		middleware: dependencies.Middleware,
	}

	return accountRouter
}

func (router *AccountRouter) RegisterRouter(engine *gin.Engine) {
	router.routerGroup = engine.Group("/account")
	router.registerRoutes()
}

func (router *AccountRouter) registerRoutes() {
	router.routerGroup.POST("/", router.middleware.ResponseHandlerMiddleware(router.controller.CreateAccount))
	router.routerGroup.POST("/token", router.middleware.ResponseHandlerMiddleware(router.controller.GetAccountToken))
}
