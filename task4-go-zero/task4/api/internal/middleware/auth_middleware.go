package middleware

import (
	"context"
	"net/http"
	"strings"
	errorx "task4-go-zero/internal/error"

	"github.com/zeromicro/go-zero/rest/httpx"
	"task4-go-zero/pkg/auth"
)

type AuthMiddleware struct {
	JwtSecret string
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		JwtSecret: jwtSecret,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			httpx.Error(w, errorx.ErrInvalidCredentials)
			return
		}

		// 解析Bearer Token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			httpx.Error(w, errorx.ErrInvalidCredentials)
			return
		}

		// 验证Token
		claims, err := auth.ParseToken(parts[1], m.JwtSecret)
		if err != nil {
			httpx.Error(w, errorx.ErrInvalidCredentials)
			return
		}

		// 将用户信息存入请求上下文
		r = r.WithContext(
			context.WithValue(
				r.Context(),
				"userID",
				claims.UserID,
			),
		)
		r = r.WithContext(
			context.WithValue(
				r.Context(),
				"userRole",
				claims.Role,
			),
		)

		next(w, r)
	}
}

type contextKey string

const (
	UserIDKey   contextKey = "userID"
	UserRoleKey contextKey = "userRole"
)
