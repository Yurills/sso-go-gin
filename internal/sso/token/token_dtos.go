package token

type TokenRequest struct {
	ClientID     string `json:"client_id" binding:"required"`
	GrantType    string `json:"grant_type" binding:"required"`
	Code 	   string `json:"code" binding:"required"`
	CodeVerifier string `json:"code_verifier" binding:"required"`
	Nonce 	  string `json:"nonce"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Nonce        string `json:"nonce,omitempty"`
}