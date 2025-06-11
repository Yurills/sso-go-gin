package service

import (
	"errors"
	"sso-go-gin/internal/sso/authorize/dtos"
	"sso-go-gin/internal/sso/authorize/repository"
	"sso-go-gin/internal/sso/models"
	"sso-go-gin/pkg/utils/randomutil"
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

	//generate csrf_token
	csrfToken, err := randomutil.GenerateRandomString(64)
	if err != nil {
		return nil, errors.New("failed to generate CSRF token")
	}

	authRequestCode := &models.AuthRequestCode{
		ID:                      uuid.New(),
		ClientID:                authClient.ID,
		ResponseType:            req.ResponseType,
		State:                   req.State,
		Nonce:                   req.Nonce,
		CodeChallenge:           req.CodeChallenge,
		CodeChallengeMethod:     req.CodeChallengeMethod,
		AuthRedirectCallbackURI: authClient.AuthRedirectCallbackURI,
		SSORedirectCallbackURI:  authClient.SSORedirectCallbackURI,
		ExpiredDatetime:         time.Now().Add(5 * time.Minute),
		CreatedDatetime:         time.Now(),
	}

	if err := s.repository.SaveCSRFToken(ctx, csrfToken, authRequestCode.ID.String(), 5*time.Minute); err != nil {
		return nil, errors.New(err.Error())
	}

	if err := s.repository.SaveRequestCode(ctx, authRequestCode); err != nil {
		return nil, errors.New("failed to save request code")
	}

	authorizeResponse := &dtos.AuthroizeResponse{
		RID:     authClient.ID.String(),
		CRSFSes: csrfToken,
	}

	return authorizeResponse, nil

}
