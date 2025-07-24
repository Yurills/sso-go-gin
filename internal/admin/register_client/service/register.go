package service

import (
	"sso-go-gin/internal/admin/models"
	"sso-go-gin/internal/admin/register_client/dtos"
	"sso-go-gin/internal/admin/register_client/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegisterService struct {
	repository *repository.RegisterRepository
}

func NewRegisterService(repo *repository.RegisterRepository) *RegisterService {
	return &RegisterService{repository: repo}
}

func (s *RegisterService) RegisterClient(ctx *gin.Context, req dtos.RegisterClient) error {
	client := &models.AuthClient{
		ID:                      uuid.New(),
		Name:                    req.Name,
		Description:             req.Description,
		ClientID:                req.ClientID,
		ClientSecret:            req.ClientSecret,
		AuthRedirectCallbackURI: req.AuthRedirectCallbackURI,
		SSORedirectCallbackURI:  req.SSORedirectCallbackURI,
		Scope:                   req.Scope,
		Active:                  req.Active,
		ConfigProfile:           []byte(req.ConfigProfile),
		PrivateKey:              req.PrivateKey,
		PublicKey:               req.PublicKey,
	}

	err := s.repository.CreateClient(ctx, client)
	if err != nil {
		return err
	}
	return nil
}
