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
    c.JSON(http.StatusOK, gin.H{"access_token": access, "refresh_token": refresh})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
    var req struct {
        RefreshToken string `json:"refresh_token"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }
    access, refresh, err := h.AuthService.Refresh(req.RefreshToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"access_token": access, "refresh_token": refresh})
}

func (h *AuthHandler) Me(c *gin.Context) {
    userID := c.GetUint("user_id")
    c.JSON(http.StatusOK, gin.H{"user_id": userID})
}