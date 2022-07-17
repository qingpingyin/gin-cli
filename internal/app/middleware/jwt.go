package middleware

import (
	"github.com/gin-gonic/gin"
	"template/pkg/auth"
)

func FrontJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(auth.Authorization)
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
