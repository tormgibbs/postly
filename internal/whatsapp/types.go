package whatsapp

type Message struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	RecipientType    string `json:"recipient_type,omitempty"`
	Type             string `json:"type"`
}

type TextMessage struct {
	Message
	Text TextBody `json:"text"`
}

type TextBody struct {
	Body string `json:"body"`
}

type InteractiveMessage struct {
	Message
	Interactive InteractiveBody `json:"interactive"`
}

type InteractiveBody struct {
	Type   string             `json:"type"`
	Body   InteractiveText    `json:"body"`
	Header *InteractiveHeader `json:"header,omitempty"`
	Action InteractiveAction  `json:"action"`
	Footer *InteractiveText   `json:"footer,omitempty"`
}

type InteractiveText struct {
	Text string `json:"text"`
}

type InteractiveHeader struct {
	Type     string       `json:"type"`
	Text     string       `json:"text,omitempty"`
	Image    *MediaObject `json:"image,omitempty"`
	Document *MediaObject `json:"document,omitempty"`
	Video    *MediaObject `json:"video,omitempty"`
}

type MediaObject struct {
	Link string `json:"link"`
}

type InteractiveAction struct {
	Name       string         `json:"name"`
	Parameters map[string]any `json:"parameters"`
}