package router

import (
	"dustwitter/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter(router *gin.Engine) {
	router.POST("/message", controller.PostMassage)

	router.GET("/top", controller.Top)
	router.GET("/home", controller.Home)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
