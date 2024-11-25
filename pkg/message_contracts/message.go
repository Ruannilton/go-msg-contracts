package messagecontracts

const (
	MessageTypeEmail = "MSG_TYPE_EMAIL"
	MessageTypeSMS   = "MSG_TYPE_SMS"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}
