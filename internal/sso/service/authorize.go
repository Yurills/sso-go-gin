package service

import (
	"errors"
	"sso-go-gin/internal/sso/dtos"
	"sso-go-gin/internal/sso/models"
	"sso-go-gin/internal/sso/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthorizeService struct {
	repository *repository.AuthorizeRepository
}

func NewAuthorizeService(repo *repository.AuthorizeRepository) *AuthorizeService {
	return &AuthorizeService{repo}
}

func (s *AuthorizeService) Authorize(ctx *gin.Context, req dtos.AuthorizeRequest) (*dtos.AuthroizeResponse, error) {
	//verify required parameter
	if req.CodeChallenge == "" || req.State == "" || req.ResponseType != "code" || req.CodeChallengeMethod != "S256" {
		return nil, errors.New("invalid request parameter")
	}

	authClient, err := s.repository.GetAuthClientByID(ctx, req.ClientID)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if !authClient.Active {
		return nil, errors.New("client not active")
	}

	authRequestCode := &models.AuthRequestCode{
		ID:                      uuid.New(),
		ClientID:                authClient.ID,
		ResponseType:            req.ResponseType,
		Scope:                   req.Scope,
		State:                   req.State,
		Nonce:                   req.Nonce,
		CodeChallenge:           req.CodeChallenge,
		CodeChallengeMethod:     req.CodeChallengeMethod,
		AuthRedirectCallbackURI: req.RedirectURI,
		SSORedirectCallbackURI:  req.RedirectURI,
		ExpiredDatetime:         time.Now().Add(10 * time.Minute),
		CreatedDatetime:         time.Now(),
	}

	if err := s.repository.SaveRequestCode(ctx, authRequestCode); err != nil {
		return nil, errors.New("failed to save request code")
	}

	authorizeResponse := &dtos.AuthroizeResponse{
		RID: authClient.ID.String(),
	}

	return authorizeResponse, nil

}
