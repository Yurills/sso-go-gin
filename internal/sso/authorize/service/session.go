package service

import (
	"github.com/gin-gonic/gin"
)

func (s *AuthorizeService) ValidSession(ctx *gin.Context, sessionID string) bool {
	session, err := s.repository.GetSessionByID(ctx, sessionID)
	if err != nil {
		return false
	}
	if session == nil || session.IsExpired() {
		return false
	}
	return true
}
