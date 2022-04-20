package middlewares

import (
	"Securing-Our-App-With-JSON-Web-Token/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		VerifyToken, err := helpers.VerifyToken(c)
		_ = VerifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", VerifyToken)
		c.Next()
	}
}
