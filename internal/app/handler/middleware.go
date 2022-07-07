package handler

import (
	"os"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func CheckKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			key := c.Request.Header.Get("Key")
			key_env, exists := os.LookupEnv("KEY")
			if exists {
				if key_env != key {
					newErrorResponse(c, 204, "Key invalid")
					return
				}
			}
		}

		c.Next()
	}
}
