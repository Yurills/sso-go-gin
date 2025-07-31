package service

import (
	"net/http"
	"sso-go-gin/internal/middleware/protect/repository"
	"sso-go-gin/internal/sso/models"

	"github.com/gin-gonic/gin"
)

type ProtectService struct {
	repository *repository.ProtectRepository
}

func NewProtectService(repository *repository.ProtectRepository) *ProtectService {
	return &ProtectService{repository: repository}
}

func (s *ProtectService) ProtectAdmin(ctx *gin.Context, sessionID string) (*models.User, error) {
	if sessionID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Session ID is required"})
		return nil, nil
	}

	userID, err := s.repository.GetUserIDBySessionID(ctx, sessionID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return nil, err
	}

	user, err := s.repository.GetUserByID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return nil, err
	}

	if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return nil, nil
	}

	if user.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return nil, nil
	}

	return user, nil
}
