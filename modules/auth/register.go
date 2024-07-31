package auth

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	db "github.com/piotroszko/backend-go/database"
	"github.com/piotroszko/backend-go/database/models"
	"github.com/piotroszko/backend-go/helpers/response"
)

type RegisterRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
type RegisterSuccess struct {
	Id    uint   `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func register(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.InvalidRequest(c)
		return
	}
	user, err := createUser(request.Login, request.Password, request.Email)
	if err != nil {
		response.InternalServerError(c)
		return
	}

	token, err := GenerateJWT(*user)
	if err != nil {
		response.InternalServerError(c)
		return
	}
	response.Ok(c, RegisterSuccess{
		Id:    user.ID,
		Login: user.Login,
		Email: user.Email,
		Token: token,
	})
}

func createUser(login, password string, email string) (*models.User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	user := &models.User{Login: login, Password: string(hashed), Email: email}

	db := db.DBConn.Create(user)
	if db.Error != nil {
		return nil, db.Error
	}

	return user, nil
}
