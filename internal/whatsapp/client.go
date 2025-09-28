package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) Send(msg any) error {
	url := fmt.Sprintf("https://graph.facebook.com/v23.0/%s/messages", c.phoneID)

	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("whatsapp api returned status: %d, response: %s", resp.StatusCode, string(respBody))
	}

	return nil
}
