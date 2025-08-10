package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, role, secret string, expiry int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expiry) * time.Second)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		logx.Errorf("Generate token error: %v", err)
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString, secret string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		logx.Errorf("Parse token error: %v", err)
		return nil, err
	}

	if !token.Valid {
		logx.Error("Invalid token")
		return nil, jwt.ErrTokenMalformed
	}

	return claims, nil
}
