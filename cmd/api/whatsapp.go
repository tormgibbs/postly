package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (app *application) handleWhatsAppWebhook(ctx *gin.Context) {
	var payload map[string]any
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	bytes, _ := json.MarshalIndent(payload, "", "  ")
	ctx.String(http.StatusOK, string(bytes))
}
