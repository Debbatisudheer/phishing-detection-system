package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(
	"phishing-platform-secret",
)

func GenerateToken(
	username string,
) (string, error) {

	token := jwtlib.NewWithClaims(
		jwtlib.SigningMethodHS256,
		jwtlib.MapClaims{
			"username": username,
			"exp": time.Now().
				Add(24 * time.Hour).
				Unix(),
		},
	)

	return token.SignedString(
		SecretKey,
	)
}