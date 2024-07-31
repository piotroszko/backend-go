package router

import (
	"github.com/gin-gonic/gin"
	"github.com/piotroszko/backend-go/modules/auth"
)

func AddV1Routes(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		auth.AddAuthGroup(v1)
	}
}
