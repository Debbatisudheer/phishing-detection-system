package models

type AnalysisResult struct {
	FileName  string
	RiskScore int
	RiskLevel string
	Verdict   string
}