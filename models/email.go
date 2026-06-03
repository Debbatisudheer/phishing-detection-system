package models

type Email struct {
	Sender      string   `json:"sender"`
	Subject     string   `json:"subject"`
	Body        string   `json:"body"`
	Attachments []string `json:"attachments"`
}