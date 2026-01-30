package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint) (string, error) {
	// payload
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // artinya setelah 24 jam token invalid
	}
	// token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// singkatnya kita merubah secret key ke byte lalu kita sign ke jwt
	return token.SignedString([]byte(os.Getenv("SECRET_KEY"))) //mengkonvert secret key ke byte/biner
}
