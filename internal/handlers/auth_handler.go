package handler

import (
	"net/http"

	"github.com/ardianilyas/go-auth/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    AuthService *services.AuthService
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }
    if err := h.AuthService.Register(req.Email, req.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "registration failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }
    access, refresh, err := h.AuthService.Login(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }
    
    c.SetCookie("access_token", access, 900, "/", "", true, true)
    c.SetCookie("refresh_token", refresh, 604800, "/", "", true, true)

    c.JSON(http.StatusOK, gin.H{"message": "login success"})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
    refreshToken, err := c.Cookie("refresh_token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
        return
    }

    access, refresh, err := h.AuthService.Refresh(refreshToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
        return
    }
    
    c.SetCookie("access_token", access, 900, "/", "", false, true)
    c.SetCookie("refresh_token", refresh, 604800, "/", "", false, true)

    c.JSON(http.StatusOK, gin.H{"message": "token refreshed"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
    c.SetCookie("access_token", "", -1, "/", "", false, true)
    c.SetCookie("refresh_token", "", -1, "/", "", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}

func (h *AuthHandler) Me(c *gin.Context) {
    userID := c.GetUint("user_id")
    c.JSON(http.StatusOK, gin.H{"user_id": userID})
}