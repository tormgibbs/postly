package main

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	cfg "github.com/tormgibbs/postly/config"
	"github.com/tormgibbs/postly/internal/db"
	"github.com/tormgibbs/postly/internal/logger"
	"github.com/tormgibbs/postly/internal/whatsapp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


type application struct {
	config   *cfg.Config
	oauth    *oauth2.Config
	queries  *db.Queries
	whatsapp *whatsapp.Client
	logger   *logger.Logger
}

func main() {
	config := cfg.LoadConfig()

	logger := logger.New(slog.LevelInfo)

	dbConn, err := openDB(*config)
	if err != nil {
		log.Fatalf("couldn't open database: %s", err.Error())
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

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
		config:   config,
		oauth:    oauth,
		queries:  queries,
		logger:   logger,
		whatsapp: whatsapp.NewClient(config.WhatsAppPhoneID, config.WhatsAppAccessToken),
	}

	server := &http.Server{
		Addr:    ":5000",
		Handler: app.routes(),
	}

	app.logger.Info("Starting server", "addr", server.Addr)

	
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		app.logger.Fatal("could not start server", "err", err)
	}

}

func openDB(c cfg.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", c.DSN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// type Message struct {
// 	ID string `json:"id"`
// }

// type Contact struct {
// 	Input string `json:"input"`
// 	WaID  string `json:"wa_id"`
// }

// type WhatsAppResponse struct {
// 	MessagingProduct string    `json:"messaging_product"`
// 	Contacts         []Contact `json:"contacts"`
// 	Messages         []Message `json:"messages"`
// }