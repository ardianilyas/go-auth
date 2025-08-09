package services

import (
	"github.com/ardianilyas/go-auth/internal/models"
	"github.com/ardianilyas/go-auth/internal/repositories"
	"github.com/ardianilyas/go-auth/pkg/token"
	"golang.org/x/crypto/bcrypt"
)
type AuthService struct {
    UserRepo *repositories.UserRepository
}

func (s *AuthService) Register(email, password string) error {
    hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    user := &model.User{Email: email, PasswordHash: string(hash)}
    return s.UserRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, string, error) {
    user, err := s.UserRepo.FindByEmail(email)
    if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
        return "", "", err
    }
    access, _ := token.GenerateAccessToken(user.ID)
    refresh, _ := token.GenerateRefreshToken(user.ID)
    s.UserRepo.UpdateRefreshToken(user.ID, refresh)
    return access, refresh, nil
}

func (s *AuthService) Refresh(oldToken string) (string, string, error) {
    claims, err := token.ParseToken(oldToken)
    if err != nil {
        return "", "", err
    }
    user, err := s.UserRepo.FindById(claims.UserID)
    if err != nil || user.RefreshToken != oldToken {
        return "", "", err
    }
    newAccess, _ := token.GenerateAccessToken(user.ID)
    newRefresh, _ := token.GenerateRefreshToken(user.ID)
    s.UserRepo.UpdateRefreshToken(user.ID, newRefresh)
    return newAccess, newRefresh, nil
}