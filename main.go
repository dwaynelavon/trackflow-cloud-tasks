package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dwaynelavon/weissach/trackflow-cloud-tasks/handlers"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

var pingPath = "/ping"

func main() {
	sentryDsn := os.Getenv("SENTRY_DSN")
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDsn,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	app := gin.Default()
	app.Use(sentrygin.New(sentrygin.Options{}))
	app.GET(pingPath, pingHandler)
	app.POST(handlers.CompleteSignUpPath, handlers.CompleteSignUpHandler)
	app.POST(handlers.SendEmailTaskHandlerPath, handlers.SendEmailHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := app.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}

/**
 * pingHandler serves as a health check.
 */
func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
