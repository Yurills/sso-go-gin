package logout

import "sso-go-gin/pkg/utils/tokenutil"

type LogoutService struct {
	repository *LogoutRepository
}

func NewLogoutService(repo *LogoutRepository) *LogoutService {
	return &LogoutService{repository: repo}
}

func (s *LogoutService) Logout(accessToken string) error {
	claims, err := tokenutil.ParseAndValidateToken(accessToken)
	if err != nil {
		return err
	}

	userID, err := s.repository.GetUserIDByEmail(claims["Email"].(string))
	if err != nil {
		return err
	}

	if err := s.repository.DeleteSessionByID(userID); err != nil {
		return err
	}
	return nil
}
