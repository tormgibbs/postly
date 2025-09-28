package whatsapp

func (c *Client) SendText(to, text string) error {
	msg := TextMessage{
		Message: Message{
			MessagingProduct: "whatsapp",
			To:               to,
			Type:             "text",
		},
		Text: TextBody{
			Body: text,
		},
	}
	return c.Send(msg)
}

func (c *Client) SendInteractive(to string, body InteractiveBody) error {
	msg := InteractiveMessage{
		Message: Message{
			MessagingProduct: "whatsapp",
			To:               to,
			Type:             "interactive",
		},
		Interactive: body,
	}
	return c.Send(msg)
}
