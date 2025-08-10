package middlewares

import (
	"net/http"

	"github.com/ardianilyas/go-auth/pkg/csrf"
	"github.com/gin-gonic/gin"
)

func CSRFMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodGet || ctx.Request.Method == http.MethodOptions {
			ctx.Next()
			return
		}

		headerToken := ctx.GetHeader("X-CSRF-TOKEN")
		cookieToken, err := ctx.Cookie("csrf_token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing CSRF token"})
			return
		}

		if err := csrf.ValidateToken(headerToken, cookieToken); err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid CSRF token"})
			return 
		}

		ctx.Next()
	}
}