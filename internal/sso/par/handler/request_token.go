package handler

import (
	"log"
	"sso-go-gin/internal/sso/par/dtos"

	"github.com/gin-gonic/gin"
)

func (h *PARHandler) PostRequestToken(c *gin.Context) {
	// Call the service to handle the request token
	// log.Printf(c.Request.Header.Get("Authorization"))
	var req dtos.PARRequestTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	//call service to handle request token
	response, err := h.Service.GetRequestToken(c, &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// c.Redirect(http.StatusFound, req.DestinationLink+"?sso_token="+response.Token)
	//TODO: construct the destination link at client NOT server side, dumbass.
	destination := response.RedirectURI + "?sso_token=" + response.Token
	log.Println(destination)
	c.JSON(200, gin.H{
		"destination_link": destination,
		"sso_token":        response.Token,
		// "token": response.Token,
	})
}
