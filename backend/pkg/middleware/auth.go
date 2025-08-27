package middleware

import (
	"backend/config"
	"backend/pkg/constant"
	"backend/pkg/httputil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			httputil.Error(ctx, http.StatusUnauthorized, "Authorization header required")
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			httputil.Error(ctx, http.StatusUnauthorized, "Authorization header format is invalid")
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return []byte(config.Cfg.Auth.SecretKey), nil
		})

		if err != nil || !token.Valid {
			httputil.Error(ctx, http.StatusUnauthorized, "Invalid or expired token")
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			httputil.Error(ctx, http.StatusUnauthorized, "Invalid token claims")
			ctx.Abort()
			return
		}

		ctx.Set("userID", int64(claims["user_id"].(float64)))
		ctx.Set("email", claims["email"])
		ctx.Set("role", claims["role"])

		ctx.Next()
	}
}

func AuthManagerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetString("role") != constant.UserRoleManager {
			httputil.Error(ctx, http.StatusForbidden, "Access denied")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
