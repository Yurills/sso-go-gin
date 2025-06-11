package service

import (
	// "sso-go-gin/internal/sso/par/dtos"
	"sso-go-gin/internal/sso/par/repository"
	// "time"

	// "sso-go-gin/internal/sso/models"

	// "errors"

	// "github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type PARService struct {
	repository *repository.PARRepository
}

func NewPARService(repository *repository.PARRepository) *PARService {
	return &PARService{repository: repository}
}

// func (s *PARService) GetRequestURI(ctx *gin.Context, req dtos.PARRequest) (string, error) {
// 	//verify required parameters
// 	if req.ClientID == "" {
// 		return "", errors.New("client ID is required")
// 	}
// 	authRequestCode, err := s.repository.GetAuthRequestByClientID(ctx, req.ClientID)
// 	if err != nil {
// 		return "", errors.New("authorization request not found for client ID: " + req.ClientID)
// 	}
// 	if authRequestCode.IsExpired() {
// 		return "", errors.New("authorization request is expired for client ID: " + req.ClientID)
// 	}
// 	authClient, err := s.repository.GetAuthClientByID(ctx, req.ClientID)
// 	if err != nil {
// 		return "", errors.New("failed to get auth client")
// 	}
// 	if !authClient.Active {
// 		return "", errors.New("auth client is not active for client ID: " + req.ClientID)
// 	}
// 	sso_token, err := s.repository.GetSSOTokenByClientID(ctx, req.ClientID)
// 	if err != nil {
// 		return "", errors.New("failed to get SSO token")
// 	}
// 	if sso_token.IsExpired() {
// 		return "", errors.New("SSO token is expired for client ID: " + req.ClientID)
// 	}

// 	if req.RedirectURI != authClient.AuthRedirectCallbackURI {
// 		return "", errors.New("redirect URI does not match the registered callback URI for client ID: " + req.ClientID)
// 	}

// 	inserted_authRequestCode := &models.AuthRequestCode{
// 		ID:                     uuid.New(),
// 		ClientID:               uuid.MustParse(req.ClientID),
// 		ResponseType:           "code",
// 		State:                  req.State,
// 		CodeChallenge:          req.CodeChallenge,
// 		CodeChallengeMethod:    req.CodeChallengeMethod,
// 		SSORedirectCallbackURI: req.RedirectURI,
// 		ExpiredDatetime:        time.Now(),
// 		CreatedDatetime:        time.Now().Add(5 * time.Minute), // Set expiration time to 5 minutes
// 	}
// 	if err := s.repository.SaveAuthRequest(ctx, inserted_authRequestCode); err != nil {
// 		return "", errors.New("failed to save auth request")
// 	}

// }
