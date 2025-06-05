package sso

type LoginRequest struct {
	Username  string `form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	CSRFToken string `form:"csrf_token" binding:"required"`
	RID       string `form:"rid" binding:"required"`
}