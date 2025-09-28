package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SendMessage(recipient, message string) error {
	accessToken := os.Getenv("WHATSAPP_ACCESS_TOKEN")
	phoneNumberID := os.Getenv("WHATSAPP_PHONE_NUMBER_ID")

	url := fmt.Sprintf("https://graph.facebook.com/v23.0/%s/messages", phoneNumberID)

	payload := map[string]any{
		"messaging_product": "whatsapp",
		"to":                recipient,
		"type":              "text",
		"text": map[string]string{
			"body": message,
		},
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	fmt.Println("WhatsApp response:", string(raw))
	return nil
}
