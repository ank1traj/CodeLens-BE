package main

import (
	"net/http"
	"os"
	"testcase/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	router.MaxMultipartMemory = 200 << 20
	router.POST("/generateInteger", handlers.GenerateRandomIntegersHandler)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
