package routers_interface

import "github.com/gin-gonic/gin"

type RouterInterface interface {
	RegisterRouter(engine *gin.Engine)
}
