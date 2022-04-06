package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", newResponse(http.StatusOK, http.StatusText(http.StatusOK), "OK"))
}
func Top(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", newResponse(http.StatusOK, http.StatusText(http.StatusOK), "OK"))
}
func PostMassage(c *gin.Context) {
	c.JSON(http.StatusCreated, newResponse(http.StatusOK, http.StatusText(http.StatusOK), "OK"))
	// Output: Json
}
