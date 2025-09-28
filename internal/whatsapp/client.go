package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	phoneID string
	token   string
}

func NewClient(phoneID, token string) *Client {
	return &Client{
		phoneID: phoneID,
		token:   token,
	}
}


func (c *Client) SendText(to, text string) error {
	url := fmt.Sprintf("https://graph.facebook.com/v18.0/%s/messages", c.phoneID)

	body := map[string]any{
		"messaging_product": "whatsapp",
		"to":                to,
		"type":              "text",
		"text": map[string]string{
			"body": text,
		},
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("whatsapp api returned %d", resp.StatusCode)
	}

	return nil
}
