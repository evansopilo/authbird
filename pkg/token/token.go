package token

import (
	"log"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID string `json:"user_id,omitempty"`
	jwt.StandardClaims
}

func Verify(tokenStr string, secret string) (*Claims, error) {

	var claims Claims

	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err)
	}

	if !token.Valid {
		return nil, err
	}

	return &claims, nil
}

func GenerateToken(claims Claims, secret string) (*string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &tokenStr, nil
}
