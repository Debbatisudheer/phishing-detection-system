package models

type EmailResponse struct {
	RiskScore int      `json:"risk_score"`
	Decision  string   `json:"decision"`
	URLs      []string `json:"urls"`
}