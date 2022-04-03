package main

import (
	"dustwitter/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	userAgent := "undefined"
	engine.LoadHTMLGlob("templates/*")

	router.SetRouter(engine)

	// ミドルウェア UserAgentの取得
	engine.Use(func(c *gin.Context) {
		userAgent = c.GetHeader("User-Agent")
		c.Next()
	})

	err := engine.Run(":3001")
	if err != nil {
		return
	}
}
