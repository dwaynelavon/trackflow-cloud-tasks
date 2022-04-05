package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dwaynelavon/weissach/trackflow-cloud-tasks/handlers"
	"github.com/gin-gonic/gin"
)

var pingPath = "/ping"

func main() {
	r := gin.Default()
	r.GET(pingPath, pingHandler)
	r.POST(handlers.CompleteSignUpPath, handlers.CompleteSignUpHandler)
	r.POST(handlers.SendEmailTaskHandlerPath, handlers.SendEmailHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
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
