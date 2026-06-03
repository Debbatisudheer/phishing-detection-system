package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/internal/parser"
	"phishing-platform/internal/domain"
	"phishing-platform/internal/threatfeed"
	"phishing-platform/internal/risk"
	"phishing-platform/internal/phishtank"
	
)

type AnalyzeEmailRequest struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type AnalyzeEmailResponse struct {
	URLs       []string `json:"urls"`
	Findings   []string `json:"findings"`
	RiskScore  int      `json:"risk_score"`
	RiskLevel  string   `json:"risk_level"`
}

func AnalyzeEmailHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req AnalyzeEmailRequest

	err :=
		json.NewDecoder(
			r.Body,
		).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"Invalid Request",
			http.StatusBadRequest,
		)

		return
	}

	urls :=
	parser.ExtractURLs(
		req.Body,
	)

	var findings []string

for _, extractedURL := range urls {

	domainFindings :=
		domain.AnalyzeURL(
			extractedURL,
		)

	findings = append(
		findings,
		domainFindings...,
	)

	threatFeedFindings :=
		threatfeed.CheckThreatFeed(
			extractedURL,
		)

	findings = append(
		findings,
		threatFeedFindings...,
	)

	phishTankFindings :=
	phishtank.CheckPhishTank(
		extractedURL,
	)

findings = append(
	findings,
	phishTankFindings...,
)
}



riskScore :=
	risk.CalculateRisk(
		req.Subject,
		req.Body,
		urls,
		findings,
	)

riskLevel :=
	risk.GetRiskLevel(
		riskScore,
	)

response :=
	AnalyzeEmailResponse{
		URLs:      urls,
		Findings:  findings,
		RiskScore: riskScore,
		RiskLevel: riskLevel,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		response,
	)
}