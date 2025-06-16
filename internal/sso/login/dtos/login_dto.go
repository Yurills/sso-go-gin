package dtos

type LoginRequest struct {
	ClientID    string `json:"client_id" binding:"required" gorm:"not null"`
	RID         string `json:"rid" binding:"required" gorm:"not null"`
	Username    string `json:"username" binding:"required" gorm:"not null"`
	Password    string `json:"password" binding:"required" gorm:"not null"`
	CSRFSession string `json:"csrf_ses" binding:"required" gorm:"not null"`
}

type LoginResponse struct {
	AuthCode    string  `json:"auth_code"`
	RedirectURI string  `json:"redirect_uri"`
	State       string  `json:"state"`
	Nonce       *string `json:"nonce,omitempty"`
}

type TokenRequest struct {
	ClientID     string `json:"client_id" binding:"required" gorm:"not null"`
	GrantType    string `json:"grant_type" binding:"required" gorm:"not null"`
	Code         string `json:"code" binding:"required" gorm:"not null"`
	CodeVerifier string `json:"code_verifier" binding:"required" gorm:"not null"`
	Nonce        string `json:"nonce"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Nonce        string `json:"nonce,omitempty"`
}

type PushAuthorizationRequest struct {
	ClientID            string `json:"client_id" binding:"required" gorm:"not null"`
	SSOToken            string `json:"sso_token" binding:"required" gorm:"not null"`
	State               string `json:"state" binding:"required" gorm:"not null"`
	CodeChallenge       string `json:"code_challenge" binding:"required" gorm:"not null"`
	CodeChallengeMethod string `json:"code_challenge_method" binding:"required" gorm:"not null"`
	RedirectURI         string `json:"redirect_uri" binding:"required" gorm:"not null"`
}
