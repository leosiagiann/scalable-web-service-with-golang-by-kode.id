package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const USERNAME = "leonardo"
const PASSWORD = "leonardo"

func Auth(c *gin.Context) bool {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "something wrong"})
		return false
	}
	if username != USERNAME || password != PASSWORD {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid usernam or password"})
		return false
	}
	return true
}
