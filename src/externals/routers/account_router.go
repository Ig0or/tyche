package routers

import "github.com/gin-gonic/gin"

type AccountRouter struct {
	routerGroup *gin.RouterGroup
}

func NewAccountRouter(engine *gin.Engine) *AccountRouter {
	accountRouter := &AccountRouter{routerGroup: engine.Group("/account")}

	return accountRouter
}

func (router *AccountRouter) RegisterRoutes() {
	router.routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
}
