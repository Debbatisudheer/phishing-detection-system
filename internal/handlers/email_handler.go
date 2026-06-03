package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"phishing-platform/database"
	"phishing-platform/internal/attachment"
	"phishing-platform/internal/decision"
	"phishing-platform/internal/domain"
	"phishing-platform/internal/parser"
	"phishing-platform/internal/risk"
	"phishing-platform/internal/threatintel"
	"phishing-platform/internal/ueba"
	"phishing-platform/models"
)

func EmailHandler(w http.ResponseWriter, r *http.Request) {

	var email models.Email

	err := json.NewDecoder(r.Body).Decode(&email)

	if err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)

		return
	}

	fmt.Println("Sender:", email.Sender)

	fmt.Println("Subject:", email.Subject)

	fmt.Println("Body:", email.Body)

	// Extract URLs
	urls := parser.ExtractURLs(email.Body)

	fmt.Println("Extracted URLs:", urls)

	var allFindings []string

	// Attachment Analysis
	attachmentFindings := attachment.AnalyzeAttachments(
		email.Attachments,
	)

	fmt.Println(
		"Attachment Findings:",
		attachmentFindings,
	)

	allFindings = append(
		allFindings,
		attachmentFindings...,
	)

	// URL + Domain + Threat Intel Analysis
	for _, extractedURL := range urls {

		findings := domain.AnalyzeURL(extractedURL)

		threatIntelFindings :=
			threatintel.CheckThreatIntel(
				extractedURL,
			)

		fmt.Println(
			"Threat Intel Findings:",
			threatIntelFindings,
		)

		findings = append(
			findings,
			threatIntelFindings...,
		)

		fmt.Println(
			"Domain Findings:",
			findings,
		)

		allFindings = append(
			allFindings,
			findings...,
		)
	}

	// UEBA Simulation
	previousLogin := ueba.LoginEvent{
		User:      "sudheer",
		Country:   "India",
		Timestamp: time.Now().Add(-30 * time.Minute),
	}

	currentLogin := ueba.LoginEvent{
		User:      "sudheer",
		Country:   "USA",
		Timestamp: time.Now(),
	}

	uebaFindings :=
		ueba.DetectImpossibleTravel(
			previousLogin,
			currentLogin,
		)

	fmt.Println(
		"UEBA Findings:",
		uebaFindings,
	)

	allFindings = append(
		allFindings,
		uebaFindings...,
	)

	// Risk Calculation
	riskScore := risk.CalculateRisk(
		email.Subject,
		email.Body,
		urls,
		allFindings,
	)

	fmt.Println("Risk Score:", riskScore)

	// Decision Engine
	decisionResult := decision.MakeDecision(
		riskScore,
	)

	fmt.Println("Decision:", decisionResult)

	// Database Insert
	query := `
INSERT INTO public.emails (
	sender,
	subject,
	body,
	risk_score,
	decision,
	findings,
	attachments
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

	_, err = database.DB.Exec(
		query,
		email.Sender,
		email.Subject,
		email.Body,
		riskScore,
		decisionResult,
		strings.Join(allFindings, ", "),
		strings.Join(email.Attachments, ", "),
	)

	if err != nil {
		fmt.Println(
			"Database Insert Error:",
			err,
		)
	}

	// API Response
	response := models.EmailResponse{
		RiskScore: riskScore,
		Decision:  decisionResult,
		URLs:      urls,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
}