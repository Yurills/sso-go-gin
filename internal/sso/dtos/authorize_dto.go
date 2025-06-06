package dtos

type AuthorizeRequest struct {
	ClientID            string `json:"client_id" binding:"required" gorm:"not null"`
	ResponseType        string `json:"response_type" binding:"required" gorm:"not null"`
	State               string `json:"state" binding:"required" gorm:"not null"`
	Scope               string `json:"scope"`
	RedirectURI         string `json:"redirect_uri" binding:"required" gorm:"not null"`
	CodeChallenge       string `json:"code_challenge" binding:"required" gorm:"not null"`
	CodeChallengeMethod string `json:"code_challenge_method" binding:"required" gorm:"not null"`
	Nonce               string `json:"nonce"`
}
