package dtos

type RegisterClient struct {
	Name                    string `json:"name" binding:"required"`
	Description             string `json:"description" binding:"required"`
	ClientID                string `json:"client_id" binding:"required"`
	ClientSecret            string `json:"client_secret" binding:"required"`
	AuthRedirectCallbackURI string `json:"auth_redirect_callback_uri" binding:"required"`
	SSORedirectCallbackURI  string `json:"sso_redirect_callback_uri" binding:"required"`
	Scope                   string `json:"scope"`
	Active                  bool   `json:"active" binding:"required"`
	ConfigProfile           string `json:"config_profile" binding:"required"`
	PrivateKey              string `json:"private_key" binding:"required"`
	PublicKey               string `json:"public_key" binding:"required"`
}
