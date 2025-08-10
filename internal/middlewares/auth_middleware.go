package middlewares

import (
	"net/http"

	"github.com/ardianilyas/go-auth/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr, err := c.Cookie("access_token")
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }

        claims, err := token.ParseToken(tokenStr)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }

        c.Set("user_id", claims.UserID)
        c.Next()
    }
}