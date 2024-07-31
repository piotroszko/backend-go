package auth

import "github.com/gin-gonic/gin"

func AddAuthGroup(g *gin.RouterGroup) {
	auth := g.Group("/auth")
	{
		auth.POST("/login", login)
		auth.POST("/register", register)
		auth.POST("/refresh", refresh)
	}
}
