package auth

import "github.com/gin-gonic/gin"

func logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}
