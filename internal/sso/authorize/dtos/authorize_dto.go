package dtos

type AuthorizeRequest struct {
	ClientID            string `form:"client_id" binding:"required" gorm:"not null"`
	ResponseType        string `form:"response_type" binding:"required" gorm:"not null"`
	State               string `form:"state" binding:"required" gorm:"not null"`
	Scope               string `form:"scope"`
	RedirectURI         string `form:"redirect_uri" binding:"required" gorm:"not null"`
	CodeChallenge       string `form:"code_challenge" binding:"required" gorm:"not null"`
	CodeChallengeMethod string `form:"code_challenge_method" binding:"required" gorm:"not null"`
	Nonce               string `form:"nonce"`
}

type AuthorizeResponse struct {
	RID     string `json:"r_id"`
	CRSFSes string `json:"csrf_ses"`
}

type AuthorizeSessionResponse struct {
	AuthCode    string  `json:"auth_code"`
	RedirectURI string  `json:"redirect_uri"`
	State       string  `json:"state"`
	Nonce       *string `json:"nonce,omitempty"`
}
