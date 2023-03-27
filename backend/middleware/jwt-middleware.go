package middleware

import (
	"fmt"
	"net/http"

	"shirt-store/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := utils.ValidateToken(c)
		fmt.Println(err)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
		}
		c.Next()
	}
}
