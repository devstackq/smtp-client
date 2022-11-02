package models

type Mail struct {
	From       string   `json:"from"`
	To         []string `json:"to"`
	Subject    string   `json:"subject"`
	Body       []byte
	Message    string `json:"message"`
	TmplTypeID string `json:"template_id"`
	DateBirth  string
}
