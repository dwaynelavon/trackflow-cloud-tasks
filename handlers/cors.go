package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type corsHeader struct {
	Origin string `header:"origin"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header corsHeader
		if err := c.BindHeader(&header); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		envAllowOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
		allowedOrigins := strings.Split(envAllowOrigins, ",")
		isOriginAllowed := false
		for _, v := range allowedOrigins {
			if v == envAllowOrigins || v == "*" {
				isOriginAllowed = true
			}
		}

		if !isOriginAllowed {
			fmt.Printf("Origin %s is not allowed.\n", header.Origin)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Writer.Header().Set(
			"Access-Control-Allow-Origin",
			header.Origin,
		)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Appengine-Taskname, X-Appengine-Queuename",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"POST, OPTIONS, GET, PUT",
		)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
