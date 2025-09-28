package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine {
	router := gin.Default()

	router.GET("/webhook", app.handleWhatsAppWebhookVerification)
	router.POST("/webhook", app.handleWhatsAppWebhook)

	router.GET("/oauth2callback", app.handleOAuth2Callback)
	router.GET("/login", app.handleGoogleLogin)

	return router
}
