package integration

import (
	"strings"
	"testing"

	"phishing-platform/internal/decision"
	"phishing-platform/internal/domain"
	"phishing-platform/internal/header"
	"phishing-platform/internal/mitre"
	"phishing-platform/internal/parser"
	"phishing-platform/internal/risk"
	"phishing-platform/internal/sender"
)

func TestCompleteEmailPipeline(t *testing.T) {

	rawEmail := `From: PayPal Support <support@evil-paypal.xyz>
Reply-To: attacker@evil.com
Return-Path: bounce@evil.com
Subject: Verify Your PayPal Account Immediately
Content-Type: text/plain

Dear Customer,

Click here immediately.

https://security-paypal.xyz/login

`

	// ----------------------------
	// Parse Email
	// ----------------------------

	email, err := parser.ParseRawEmail(
		strings.NewReader(rawEmail),
	)

	if err != nil {
		t.Fatal(err)
	}

	// ----------------------------
	// URLs
	// ----------------------------

	urls := parser.ExtractURLs(
		email.Body,
	)

	// ----------------------------
	// Header Analysis
	// ----------------------------

	var findings []string

	findings = append(
		findings,
		header.AnalyzeHeaders(
			email.From,
			email.ReplyTo,
			email.ReturnPath,
		)...,
	)

	findings = append(
		findings,
		header.DetectDisplayNameSpoofing(
			email.From,
		)...,
	)

	// ----------------------------
	// Sender Reputation
	// ----------------------------

	findings = append(
		findings,
		sender.CheckSenderReputation(
			email.From,
		)...,
	)

	// ----------------------------
	// URL Analysis
	// ----------------------------

	for _, url := range urls {

		findings = append(
			findings,
			domain.AnalyzeURL(url)...,
		)
	}

	// ----------------------------
	// MITRE
	// ----------------------------

	technique := mitre.MapTechnique(
		email.Subject,
		email.Body,
	)

	if technique == "" {
		t.Fatal("expected MITRE technique")
	}

	// ----------------------------
	// Risk
	// ----------------------------

	score := risk.CalculateRisk(
		email.Subject,
		email.Body,
		urls,
		findings,
	)

	if score <= 0 {
		t.Fatal("expected positive risk score")
	}

	// ----------------------------
	// Decision
	// ----------------------------

	decisionResult := decision.MakeDecision(
		score,
	)

	if decisionResult == "" {
		t.Fatal("expected decision")
	}

	// ----------------------------
	// Final Assertions
	// ----------------------------

	if len(findings) == 0 {
		t.Fatal("expected findings")
	}

	t.Log("========== COMPLETE PIPELINE ==========")

	t.Log("Parsed Sender:", email.From)

	t.Log("URLs:", urls)

	t.Log("Findings:", findings)

	t.Log("MITRE:", technique)

	t.Log("Risk Score:", score)

	t.Log("Decision:", decisionResult)

	t.Log("=======================================")
}