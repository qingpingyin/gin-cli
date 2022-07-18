package middleware

import (
	"github.com/gin-gonic/gin"
)

func FrontJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("")
		if token == "" {
			c.Abort()
			return
		}
	}
}

func MerchantJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
