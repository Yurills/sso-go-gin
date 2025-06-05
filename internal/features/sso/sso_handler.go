package sso

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(Service *Service) *Handler {
	return &Handler{Service}
}

func (h *Handler) PostLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form submission"})
		return
	}

	//need check csrf token

	//check auth request by RID
	authReq, err := h.Service.repository.GetAuthRequestByID(c.Request.Context(), req.RID)
	if err != nil || authReq.IsExpired() {
		c.JSON(400, gin.H{"error": err})
		return
	}

	//check username password
	user, err := h.Service.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	//generate token
	_, redirectURI, err := h.Service.IssueAuthCode(c.Request.Context(), user.ID, req.RID)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"redirectedURI": redirectURI})

}
