package logout

type LogoutRequest struct {
	Username string `json:"username" binding:"required"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}
