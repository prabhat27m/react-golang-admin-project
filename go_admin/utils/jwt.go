package utils

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go/v4"
)
const SecretKey = "secret"

func GenerateJwtToken(Issuer string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour) // Adjust the expiration time as needed

    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
        Issuer:    fmt.Sprint(Issuer),
        ExpiresAt: jwt.At(expirationTime),
		
    })

    token, err := claims.SignedString([]byte(SecretKey))
    return token, err
}

func ParseJwtToken(cookie string) (string, error) {
    token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
        return []byte(SecretKey), nil
    })

    if err != nil || !token.Valid {
        return "Invalid Token", err
    }

    claims, ok := token.Claims.(*jwt.StandardClaims)
    if !ok {
        return "Invalid Claims", fmt.Errorf("unable to extract claims from token")
    }

    fmt.Println("Token expired at:", claims.ExpiresAt)

    return claims.Issuer, nil
}
