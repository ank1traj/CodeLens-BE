package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testcase/internal/handlers"
	"time"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv(key)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, baggage, sentry-trace") // Include sentry-trace header
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	DSN := goDotEnvVariable("SENTRY_DSN")
	ENVIRONMENT := goDotEnvVariable("SENTRY_ENV")

	if ENVIRONMENT == "" {
		ENVIRONMENT = "development"
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              DSN,
		Environment:      ENVIRONMENT,
		Debug:            false,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
		TracesSampler: func(ctx sentry.SamplingContext) float64 {
			return 1.0
		},
		ProfilesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	if goDotEnvVariable("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/generateInteger", handlers.GenerateRandomIntegersHandler)
	router.POST("/generateString", handlers.GenerateRandomStringsHandler)
	router.POST("/generateArray", handlers.GenerateRandomArrayHandler)

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
