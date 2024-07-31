package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/piotroszko/backend-go/helpers/response"
)

type RefreshSuccess struct {
	Token string `json:"token"`
}

func refresh(c *gin.Context) {
	user, err := GetAuthContext(c)
	if err != nil {
		response.Unauthorized(c)
		return
	}
	newToken, err := RefreshToken(user.token)
	if err != nil {
		response.Unauthorized(c)
		return
	}

	response.Ok(c, RefreshSuccess{Token: newToken})
}
