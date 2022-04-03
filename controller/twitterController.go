package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Top(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Hello Golang HTML",
	})
}
func PostMassage(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
