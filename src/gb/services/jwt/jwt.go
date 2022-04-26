package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	controller "insanitygaming.net/bans/src/gb/controllers/admin"
	"insanitygaming.net/bans/src/gb/models/admin"
)

var jwt_secret = []byte("secret")

func Build() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":       "bar",
		"ExpiresAt": time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(jwt_secret)
}

func Decode(app *gin.Context, tokenStirng string) (*admin.Admin, error) {
	token, err := jwt.Parse(tokenStirng, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwt_secret, nil
	})

	var admin *admin.Admin

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		admin, err = controller.Find(app, claims["id"].(uint))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Invalid token")
	}
	return admin, nil
}
