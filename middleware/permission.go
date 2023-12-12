package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginUser := c.MustGet(gin.AuthUserKey).(string)
		fmt.Println(loginUser)
	}
}
