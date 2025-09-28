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

	user, err := app.queries.GetUserByWhatsAppID(c, from)
	if err != nil {
		user, err = app.queries.CreateUser(c, from)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

	}

	if !user.IsLoggedIn {
		state := fmt.Sprintf("wa_%s", from)

		loginURL := app.oauth.AuthCodeURL(
			state,
			oauth2.AccessTypeOffline,
		)

		body := whatsapp.InteractiveBody{
			Type: "cta_url",
			Body: whatsapp.InteractiveText{
				Text: "👋 Hey there! Welcome aboard.\n\nTo get started, you’ll need to connect your Gmail account so we can sync your emails with WhatsApp. Don’t worry — it’s quick and secure.",
			},
			Action: whatsapp.InteractiveAction{
				Name: "cta_url",
				Parameters: map[string]any{
					"display_text": "Connect Google",
					"url":          loginURL,
				},
			},
			Footer: &whatsapp.InteractiveText{
				Text: "We’ll sync your Gmail once you approve access.",
			},
		}

		err = app.whatsapp.SendInteractive(from, body)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed send message"})
			return
		}
	}

	c.Status(http.StatusOK)
}
