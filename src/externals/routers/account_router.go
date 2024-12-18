package routers

import (
	"github.com/Ig0or/tyche/src/adapters/ports/controllers_interface"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type AccountRouter struct {
	routerGroup *gin.RouterGroup
	controller  controllers_interface.AccountControllerInterface
}

type AccountRouterDependencies struct {
	dig.In

	Controller controllers_interface.AccountControllerInterface `name:"AccountController"`
}

func NewAccountRouter(dependencies AccountRouterDependencies) *AccountRouter {
	accountRouter := &AccountRouter{controller: dependencies.Controller}

	return accountRouter
}

func (router *AccountRouter) RegisterRouter(engine *gin.Engine) {
	router.routerGroup = engine.Group("/account")
	router.registerRoutes()
}

func (router *AccountRouter) registerRoutes() {
	router.routerGroup.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
}
