package token

type TokenRequest struct {
	ClientID     string `form:"client_id" json:"client_id" binding:"required"`
	GrantType    string `form:"grant_type" json:"grant_type" binding:"required"`
	Code         string `form:"code" json:"code" binding:"required"`
	CodeVerifier string `form:"code_verifier" json:"code_verifier"`
	Nonce        string `form:"nonce" json:"nonce"`
}

type TokenResponse struct {
	AccessToken     string `json:"access_token"`
	TokenType       string `json:"token_type"`
	ExpiresIn       int    `json:"expires_in"`
	RefreshToken    string `json:"refresh_token,omitempty"`
	Nonce           string `json:"nonce,omitempty"`
	DestinationLink string `json:"destination_link,omitempty"` // Added to match the response structure
}
