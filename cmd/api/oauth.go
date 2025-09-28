package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (app *application) handleGoogleLogin(c *gin.Context) {
	url := app.oauth.AuthCodeURL("random-state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (app *application) handleOAuth2Callback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	if code != "" || state == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "missing code or state",
		})
		return
	}

	token, err := app.oauth.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "failed to exchange token",
		})
		return
	}

	var whatsappID string
	if strings.HasPrefix(state, "wa_") {
		whatsappID = strings.TrimPrefix(state, "wa_")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid state"})
		return
	}

	if err := app.queries.SetUserLoggedIn(c, whatsappID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user login"})
		return
	}

	fmt.Printf("Access token: %s\n", token.AccessToken)
	fmt.Printf("Refresh token: %s\n", token.RefreshToken)

	c.String(http.StatusOK, "Gmail connected! You can close this tab.")

	c.JSON(http.StatusOK, gin.H{
		"message":     "Login successful",
		"whatsapp_id": whatsappID,
	})

}
