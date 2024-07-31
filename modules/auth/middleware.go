package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProtectedRoute(c *gin.Context) {
	fmt.Println("Im a dummy!")

	// Pass on to the next-in-chain
	c.Next()
}
