package auth

import (
	"github.com/gin-gonic/gin"
	db "github.com/piotroszko/backend-go/database"
	"github.com/piotroszko/backend-go/database/models"
	"github.com/piotroszko/backend-go/helpers/response"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginSuccess struct {
	Id    uint   `json:"id"`
	Login string `json:"login"`
	Token string `json:"token"`
}

func login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.InvalidRequest(c)
		return
	}
	user, err := findUserAndCheckPassword(request.Login, request.Password)
	if err != nil {
		response.WrongCredentials(c)
		return
	}
	jwtToken, err := GenerateJWT(*user)
	if err != nil {
		response.InternalServerError(c)
		return
	}
	response.Ok(c, LoginSuccess{
		Id:    user.ID,
		Login: user.Login,
		Token: jwtToken,
	})
}

func findUserAndCheckPassword(login, password string) (*models.User, error) {
	var user models.User
	db := db.DBConn.Where("login = ?", login).First(&user)
	if db.Error != nil {
		return nil, db.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}
