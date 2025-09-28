package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (app *application) handleGoogleLogin(c *gin.Context) {
	url := app.oauth.AuthCodeURL("random-state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (app *application) handleOAuth2Callback(c *gin.Context) {
	code := c.Query("code")
	if code != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "missing code",
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

	fmt.Printf("Access token: %s\n", token.AccessToken)
	fmt.Printf("Refresh token: %s\n", token.RefreshToken)

	c.String(http.StatusOK, "Gmail connected! You can close this tab.")

}
