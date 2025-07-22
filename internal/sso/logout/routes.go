package logout

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, h *LogoutHandler) {
	rg.POST("/logout", h.PostLogout)
}