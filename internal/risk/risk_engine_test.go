package risk

import "testing"

func TestGetRiskLevel(t *testing.T) {

	tests := []struct {
		name     string
		score    int
		expected string
	}{
		{
			name:     "Low",
			score:    50,
			expected: "LOW",
		},
		{
			name:     "Medium",
			score:    150,
			expected: "MEDIUM",
		},
		{
			name:     "High",
			score:    350,
			expected: "HIGH",
		},
		{
			name:     "Critical",
			score:    600,
			expected: "CRITICAL",
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := GetRiskLevel(tc.score)

			if result != tc.expected {
				t.Errorf(
					"expected %s got %s",
					tc.expected,
					result,
				)
			}
		})
	}
}

func TestCalculateRisk(t *testing.T) {

	tests := []struct {
		name         string
		subject      string
		body         string
		urls         []string
		findings     []string
		expectedRisk int
	}{
		{
			name:         "Safe Email",
			subject:      "Hello",
			body:         "How are you?",
			urls:         []string{},
			findings:     []string{},
			expectedRisk: 0,
		},
		{
			name:    "URL Only",
			subject: "Hello",
			body:    "Visit our website",
			urls: []string{
				"https://example.com",
			},
			findings: []string{},
			expectedRisk: 40,
		},
		{
			name:    "Urgent Login",
			subject: "Urgent",
			body:    "Please login",
			urls:    []string{},
			findings: []string{},
			expectedRisk: 50, // 30 + 20
		},
		{
			name:    "Reply-To Mismatch",
			subject: "Hello",
			body:    "Test",
			urls:    []string{},
			findings: []string{
				"Reply-To mismatch detected",
			},
			expectedRisk: 100,
		},
		{
			name:    "Display Name Spoof",
			subject: "Hello",
			body:    "Test",
			urls:    []string{},
			findings: []string{
				"Display name spoofing detected",
			},
			expectedRisk: 120,
		},
		{
			name:    "Threat Feed Hit",
			subject: "Hello",
			body:    "Test",
			urls:    []string{},
			findings: []string{
				"Threat feed hit",
			},
			expectedRisk: 200,
		},
		{
			name:    "VirusTotal Malicious",
			subject: "Hello",
			body:    "Test",
			urls:    []string{},
			findings: []string{
				"VirusTotal malicious URL detected",
			},
			expectedRisk: 300,
		},
		{
			name:    "Critical Score Cap",
			subject: "Urgent Security Alert",
			body:    "login verify account click here password bank",
			urls: []string{
				"https://evil.com",
			},
			findings: []string{
				"Behavior Rule",
				"VirusTotal malicious URL detected",
				"Threat feed hit",
				"QR code phishing detected",
				"Persistence Detected",
				"Display name spoofing detected",
				"BEC indicator detected",
				"PhishTank hit:",
			},
			expectedRisk: 1000,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			score := CalculateRisk(
				tc.subject,
				tc.body,
				tc.urls,
				tc.findings,
			)

			if score != tc.expectedRisk {
				t.Errorf(
					"expected %d got %d",
					tc.expectedRisk,
					score,
				)
			}
		})
	}
}