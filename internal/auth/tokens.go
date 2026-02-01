package auth

import (
	"fmt"
	"strings"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractAndParseToken(auth string) (*Claims, error) {
	token, err := ExtractToken(auth)
	if err != nil {
		return nil, fmt.Errorf("ExtractAndParse: extract: %w", err)
	}
	claims, err := ParseToken(token, []byte("todo! create env and set secret"))
	if err != nil {
		return nil, fmt.Errorf("ExtractAndParse: parse: %w", err)
	}
	return claims, nil
}

func ExtractToken(auth string) (string, error) {
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return "", fmt.Errorf("invalid Authorization header format")
		}

		return parts[1], nil
}


func ParseToken(tokenStr string, secret []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
