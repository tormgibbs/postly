package main

import (
	"log"
	"net/http"

	cfg "github.com/tormgibbs/postly/config"
	"github.com/tormgibbs/postly/internal/sqlc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Message struct {
	ID string `json:"id"`
}

type Contact struct {
	Input string `json:"input"`
	WaID  string `json:"wa_id"`
}

type WhatsAppResponse struct {
	MessagingProduct string    `json:"messaging_product"`
	Contacts         []Contact `json:"contacts"`
	Messages         []Message `json:"messages"`
}

type application struct {
	config *cfg.Config
	oauth  *oauth2.Config
	store  *sqlc.Queries
}

func main() {
	config := cfg.LoadConfig()

	oauth := &oauth2.Config{
		ClientID:     config.GoogleClientID,
		ClientSecret: config.GoogleClientSecret,
		RedirectURL:  config.GoogleRedirectURI,
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}

	app := &application{
		config: config,
		oauth:  oauth,
	}

	server := &http.Server{
		Addr:    ":5000",
		Handler: app.routes(),
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not start server: %v", err)
	}

}
