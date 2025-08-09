package middlewares

import (
	"net/http"

	"github.com/ardianilyas/go-auth/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if len(auth) < 7 || auth[:7] != "Bearer " {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            return
        }
        claims, err := token.ParseToken(auth[7:])
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }
        c.Set("user_id", claims.UserID)
        c.Next()
    }
}