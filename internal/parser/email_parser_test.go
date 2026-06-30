package parser

import (
	"strings"
	"testing"
)

func TestParseRawEmail(t *testing.T) {

	rawEmail := `From: alice@example.com
Reply-To: reply@example.com
Return-Path: bounce@example.com
Subject: Test Email
Content-Type: text/plain

Hello Sudheer,
This is a test email.
Visit https://example.com
`

	email, err := ParseRawEmail(
		strings.NewReader(rawEmail),
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if email.From != "alice@example.com" {
		t.Errorf(
			"expected From alice@example.com got %s",
			email.From,
		)
	}

	if email.ReplyTo != "reply@example.com" {
		t.Errorf(
			"expected Reply-To reply@example.com got %s",
			email.ReplyTo,
		)
	}

	if email.ReturnPath != "bounce@example.com" {
		t.Errorf(
			"expected Return-Path bounce@example.com got %s",
			email.ReturnPath,
		)
	}

	if email.Subject != "Test Email" {
		t.Errorf(
			"expected Subject Test Email got %s",
			email.Subject,
		)
	}

	if !strings.Contains(
		email.Body,
		"Hello Sudheer",
	) {
		t.Error(
			"body not parsed correctly",
		)
	}

	if len(email.Attachments) != 0 {
		t.Error(
			"expected no attachments",
		)
	}
}

func TestParseRawEmail_Empty(t *testing.T) {

	_, err := ParseRawEmail(
		strings.NewReader(""),
	)

	if err == nil {
		t.Error("expected error for empty email")
	}
}

func TestParseRawEmail_Invalid(t *testing.T) {

	_, err := ParseRawEmail(
		strings.NewReader("abcdefg"),
	)

	if err == nil {
		t.Error("expected parsing error")
	}
}