package dtos

type PARRequestTokenRequest struct {
	ClientID        string `json:"client_id" binding:"required"`
	Source          string `json:"source" binding:"required"`
	Destination     string `json:"destination" binding:"required"`
	DestinationLink string `json:"destination_link" binding:"required"`
}

type PARRequestTokenResponse struct {
	Token string `json:"token" binding:"required"`
}

type PARRequest struct {
	ClientID            string `json:"client_id" binding:"required"`
	SSOToken            string `json:"sso_token" binding:"required"`
	State               string `json:"state" binding:"required"`
	CodeChallenge       string `json:"code_challenge" binding:"required"`
	CodeChallengeMethod string `json:"code_challenge_method" binding:"required"`
	RedirectURI         string `json:"redirect_uri" binding:"required"`
}

type PARResponse struct {
	RequestURI string `json:"request_uri" binding:"required"`
}

type PARRequestAuthorize struct {
	RequestURI string `json:"request_uri" binding:"required"`
}

type PARResponseAuthorize struct {
	RedirectURI string `json:"redirect_uri" binding:"required"`
	Code        string `json:"code" binding:"required"`
}
