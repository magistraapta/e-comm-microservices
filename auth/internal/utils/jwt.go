package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	UserID int64
	jwt.RegisteredClaims
}

func GenerateToken(userID int64) (string, error) {
	expirationDate := time.Now().Add(24 * time.Hour)
	claims := &JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationDate),
			IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(os.Getenv("SECRET_KEY"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Validate(signedToken string) (string, error) {
	claims := &JwtClaims{}

	token, err := jwt.ParseWithClaims(signedToken, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("Invalid token")
	}

	return claims.ID, nil
}
