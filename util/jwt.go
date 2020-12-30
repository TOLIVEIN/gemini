package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

//Claims ...
type Claims struct {
	Username    string
	Permissions string
	jwt.StandardClaims
}

//GenerateToken ...
func GenerateToken(username string) (string, error) {
	generateTime := time.Now()
	expireTime := generateTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		"user",
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gemini",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//ParseToken ...
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// fmt.Println(jwtSecret)
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
