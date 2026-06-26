package models

type SandboxReport struct {
	ID         int    `json:"id"`
	JobID      int      `json:"job_id"`
	FileName   string `json:"file_name"`
	FileSize   int64  `json:"file_size"`
	Extension  string `json:"extension"`
	MimeType   string `json:"mime_type"`
	MD5        string `json:"md5"`
	SHA256     string `json:"sha256"`
	Findings   string `json:"findings"`
	RiskScore  int    `json:"risk_score"`
	RiskLevel  string `json:"risk_level"`
	Verdict    string `json:"verdict"`
	MITRE      string `json:"mitre"`
}