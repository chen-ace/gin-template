package router

import (
	"gin-template/controller"
	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")
	{
		apiRouter.POST("/login/account", controller.Login)
		apiRouter.GET("/currentUser", controller.CurrentUser)
	}
}
