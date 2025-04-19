package utils

import (
	"auth/internal/model"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// var jwtSecretKey = []byte("nsusbfuwebfefmeinfe") // Ideally from env/config

// func GenerateToken(user model.User) (string, error) {
// 	// Create JWT claims
// 	claims := jwt.MapClaims{
// 		"user_id": user.Id,
// 		"email":   user.Email,
// 		"exp":     time.Now().Add(time.Hour * 72).Unix(), // token expires in 72 hours
// 		"iat":     time.Now().Unix(),
// 	}

// 	// Create the token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Sign the token with your secret key
// 	signedToken, err := token.SignedString(jwtSecretKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedToken, nil
// }

type JwtClaims struct {
	Id       int64
	Username string
	Email    string
	Password string
	jwt.StandardClaims
}

func GenerateToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "self-pickup",
			ExpiresAt: expirationTime.Unix(),
		},
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return "", fmt.Errorf("SECRET_KEY environment variable not set")
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateAdminToken(admin *model.Admin) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "self-pickup",
			ExpiresAt: expirationTime.Unix(),
		},
		Id:       admin.Id,
		Username: admin.Username,
		Password: admin.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return "", fmt.Errorf("SECRET_KEY environment variable not set")
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// func Validate(signedToken string) (*JwtClaims, error) {
// 	secret := os.Getenv("SECRET_KEY")
// 	if secret == "" {
// 		return nil, fmt.Errorf("SECRET_KEY environment variable not set")
// 	}

// 	token, err := jwt.ParseWithClaims(signedToken, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return []byte(secret), nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	claims, ok := token.Claims.(*JwtClaims)
// 	if !ok || !token.Valid {
// 		return nil, errors.New("invalid token")
// 	}

// 	return claims, nil
// }

func ValidateToken(signedToken string) (*JwtClaims, error) {
	secret := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(signedToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, err
}
