/*
main package
*/
package main

import (
	"dustwitter/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main function
func main() {
	engine := gin.Default()

	// godotenv load
	envLoadError := godotenv.Load(".env")
	if envLoadError != nil {
		panic("Error loading .env file")
	}

	//userAgent := "undefined"
	engine.LoadHTMLGlob("templates/*")

	router.SetRouter(engine)

	// ミドルウェア UserAgentの取得
	//engine.Use(func(c *gin.Context) {
	//	userAgent = c.GetHeader("User-Agent")
	//	c.Next()
	//})

	serverRunError := engine.Run(":3001")
	if serverRunError != nil {
		return
	}
}
