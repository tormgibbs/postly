package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tormgibbs/postly/internal/db"
	"golang.org/x/oauth2"
)

func (app *application) handleGoogleLogin(c *gin.Context) {
	url := app.oauth.AuthCodeURL("random-state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (app *application) handleOAuth2Callback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	if code == "" || state == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "missing code or state",
		})
		return
	}

	waID := strings.TrimPrefix(state, "wa_")

	user, err := app.queries.GetUserByWhatsAppID(c, waID)
	if err != nil {
		fmt.Printf("\nuser not found: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if user.IsLoggedIn {
		c.Redirect(http.StatusSeeOther, "/success")
		return
	}

	token, err := app.oauth.Exchange(c, code)
	if err != nil {
		fmt.Printf("\nfailed to exchange token: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "failed to exchange token",
		})
		return
	}

	params := db.SetUserLoggedInWithTokenParams{
		WhatsappID:   waID,
		AccessToken:  db.ToNullString(token.AccessToken),
		RefreshToken: db.ToNullString(token.RefreshToken),
		TokenExpiry:  db.ToNullTime(token.Expiry),
	}

	err = app.queries.SetUserLoggedInWithToken(c, params)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user login"})
		return
	}


	c.Redirect(http.StatusSeeOther, "/success")
}
