package middleware

import (
	"fmt"
	"github.com/bqxtt/book_online/api/auth"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims, err := auth.ParseToken(token)
		if err != nil {
			c.Next()
			return
		}
		c.Set("auth", claims)
		fmt.Println(claims)
		c.Next()
	}
}
