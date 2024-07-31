package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/piotroszko/backend-go/helpers/response"
)

type AuthContext struct {
	id    uint
	login string
	token string
}

func ProtectedRoute(c *gin.Context) {
	token := c.GetHeader("Authorization")

	id, login, err := GetIdAndLoginFromJWT(token)
	if err != nil {
		response.Unauthorized(c)
		return
	}

	c.Set("id", id)
	c.Set("login", login)
	c.Set("token", token)

	c.Next()
}

func GetAuthContext(c *gin.Context) (AuthContext, error) {
	id, errId := c.Get("id")
	login, errLogin := c.Get("login")
	token, errToken := c.Get("token")
	if !errId || !errLogin || !errToken {
		return AuthContext{}, fmt.Errorf("error getting auth context")
	}

	return AuthContext{id: id.(uint), login: login.(string), token: token.(string)}, nil
}
