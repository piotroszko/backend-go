package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/piotroszko/backend-go/database/models"
)

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, getClaims(user.ID, user.Login))
	signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func GetIdAndLoginFromJWT(tokenString string) (uint, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, "", err
	}

	if !getIsParsedTokenValid(token) {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", err
	}
	id := uint(claims["id"].(uint))
	login := claims["login"].(string)
	return id, login, nil
}

func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", err
	}

	if !getIsParsedTokenValid(token) {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, getClaims(uint(claims["id"].(uint)), claims["login"].(string)))
	signed, err := newToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func GetIsTokenValid(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false
	}

	if !getIsParsedTokenValid(token) {
		return false
	}

	return true
}

func getClaims(id uint, login string) jwt.Claims {
	return jwt.MapClaims{
		"id":    id,
		"login": login,
		"iss":   "backend-go",
		"aud":   "frontend-go",
		"nbf":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
}

func getIsParsedTokenValid(token *jwt.Token) bool {
	if !token.Valid {
		return false
	}
	date, err := token.Claims.GetExpirationTime()
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}

	return true
}
