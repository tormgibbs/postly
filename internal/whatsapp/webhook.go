package whatsapp

type WebhookPayload struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	ID      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string    `json:"messaging_product"`
	Metadata         Metadata  `json:"metadata"`
	Contacts         []Contact `json:"contacts,omitempty"`
	Messages         []Message `json:"messages,omitempty"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type Contact struct {
	Profile Profile `json:"profile"`
	WAID    string  `json:"wa_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type Message struct {
	Context   *MessageContext `json:"context,omitempty"`
	From      string          `json:"from"`
	ID        string          `json:"id"`
	Timestamp string          `json:"timestamp"`
	Type      string          `json:"type"`
	Text      *TextMessage    `json:"text,omitempty"`
	Image     *MediaMessage   `json:"image,omitempty"`
	Audio     *MediaMessage   `json:"audio,omitempty"`
	Document  *MediaMessage   `json:"document,omitempty"`
	Video     *MediaMessage   `json:"video,omitempty"`
}

type MessageContext struct {
	ID string `json:"id"`
}

type TextMessage struct {
	Body string `json:"body"`
}

type MediaMessage struct {
	ID   string `json:"id"`
	Mime string `json:"mime_type"`
}
