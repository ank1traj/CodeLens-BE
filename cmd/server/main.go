package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"testcase/internal/handlers"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router.MaxMultipartMemory = 200 << 20
	router.POST("/generateInteger", handlers.GenerateRandomIntegersHandler)
	router.POST("/generateString", handlers.GenerateRandomStringsHandler)

	err := router.Run(":8080")
	if err != nil {
		println(err)
		panic(err)
	}
}
