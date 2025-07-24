package handler

import (
	"net/http"
	"sso-go-gin/internal/admin/register_client/dtos"
	"sso-go-gin/internal/admin/register_client/service"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	service *service.RegisterService
}

func NewRegisterHandler(service *service.RegisterService) *RegisterHandler {
	return &RegisterHandler{service: service}
}

func (h *RegisterHandler) RegisterClient(ctx *gin.Context) {
	var req dtos.RegisterClient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.RegisterClient(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Client registered successfully"})
}
