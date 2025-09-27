package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

func main() {
	godotenv.Load()

	accessToken := os.Getenv("WHATSAPP_ACCESS_TOKEN")
	phoneNumberID := os.Getenv("WHATSAPP_PHONE_NUMBER_ID")
	recipient := os.Getenv("WHATSAPP_TEST_RECIPIENT")


	url := fmt.Sprintf("https://graph.facebook.com/v23.0/%s/messages", phoneNumberID)

	payload := map[string]any{
		"messaging_product": "whatsapp",
		"to":                recipient,
		"type":              "text",
		"text": map[string]string{
			"body": "hello world",
		},
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	var pretty bytes.Buffer
	json.Indent(&pretty, raw, "", "  ")

	fmt.Println("Status:", resp.Status, "\nBody:", pretty.String())
}
