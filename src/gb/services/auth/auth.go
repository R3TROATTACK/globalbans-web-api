package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Middleware(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		c.AbortWithStatus(401)
		return
	}

	admin, err := Decode(c, header[7:])
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	c.Set("user", admin)
	c.Next()
}

var jwt_secret = []byte("secret")

func Build() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":       "bar",
		"ExpiresAt": time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(jwt_secret)
}

func Decode(app *gin.Context, tokenStirng string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStirng, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwt_secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
