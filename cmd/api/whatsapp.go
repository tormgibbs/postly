package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tormgibbs/postly/internal/whatsapp"
	"golang.org/x/oauth2"
)

func (app *application) handleWhatsAppWebhookVerification(c *gin.Context) {
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	if mode == "subscribe" && token == app.config.WhatsAppWebhookToken {
		c.String(http.StatusOK, challenge)
		return
	}

	c.JSON(http.StatusForbidden, gin.H{"error": "verification failed"})
}

func (app *application) handleWhatsAppWebhook(c *gin.Context) {
	var payload whatsapp.WebhookPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	
	app.logger.Info("Received WhatsApp webhook", "payload", payload)

	if len(payload.Entry) == 0 || len(payload.Entry[0].Changes) == 0 {
		c.Status(http.StatusOK)
		return
	}

	value := payload.Entry[0].Changes[0].Value
	if len(value.Messages) == 0 {
		c.Status(http.StatusOK)
		return
	}

	msg := value.Messages[0]
	from := msg.From
	// text := ""
	// if msg.Text != nil {
	// 	text = msg.Text.Body
	// }

	user, err := app.queries.GetUserByWhatsAppID(c, from)
	if err != nil {
		user, err = app.queries.CreateUser(c, from)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

	}

	if !user.IsLoggedIn {
		// put the WhatsApp user ID in the state param
		state := fmt.Sprintf("wa_%s", from)

		loginURL := app.oauth.AuthCodeURL(
			state,
			oauth2.AccessTypeOffline,
		)

		_ = app.whatsapp.SendText(from, "Welcome! Please login: "+loginURL)
	}

	// if !user.IsLoggedIn {
	// 	loginURL := app.oauth.AuthCodeURL("state", oauth2.AccessTypeOffline)
	// 	_ = app.whatsapp.SendText(from, "Welcome! Please login: "+loginURL)
	// }

	c.Status(http.StatusOK)
}
