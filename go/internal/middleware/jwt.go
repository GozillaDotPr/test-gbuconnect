package middleware

import (
	"net/http"
	"strings"

	"github.com/app/gin-postgres-api/internal/service"
	"github.com/app/gin-postgres-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func JWTAuth(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, "missing or invalid authorization header")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := authService.ValidateToken(tokenStr)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("username", claims["username"])
		c.Next()
	}
}
