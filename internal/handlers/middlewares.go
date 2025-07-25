package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func recoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Caught the panic", err)
				respError(c, http.StatusInternalServerError, "Internal server error")
				c.Abort()
			}
		}()

		c.Next()
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		r := c.Request
		w := c.Writer

		if strings.HasPrefix(r.URL.Path, "/auth") {
			c.Next()
			return
		}

		authHeader := r.Header.Get("Authorization")
		var token string
		if parts := strings.Split(authHeader, "Bearer "); len(parts) == 2 {
			token = parts[1]
		}

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//TODO: JWT auth here

		c.Next()
	}
}
