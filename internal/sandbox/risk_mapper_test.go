package sandbox

import "testing"

func TestCalculateSandboxRisk(t *testing.T) {

	tests := []struct {
		name          string
		findings      []string
		expectedScore int
		expectedLevel string
		expectedVerd  string
	}{
		{
			name: "Low",
			findings: []string{
				"Safe",
			},
			expectedScore: 0,
			expectedLevel: "LOW",
			expectedVerd:  "ALLOW",
		},
		{
			name: "Medium",
			findings: []string{
				"Sandbox IOC URL: https://evil.com",
			},
			expectedScore: 120,
			expectedLevel: "MEDIUM",
			expectedVerd:  "SUSPICIOUS",
		},
		{
			name: "High",
			findings: []string{
				"Behavior Rule: PowerShell Downloader Detected",
				"YARA rule matched: PowerShell",
			},
			expectedScore: 600,
expectedLevel: "CRITICAL",
			expectedVerd:  "QUARANTINE",
		},
		{
			name: "Critical",
			findings: []string{
				"Behavior Rule: PowerShell Downloader Detected",
				"Persistence Detected: Registry Run Key Modification",
				"VirusTotal malicious hash detected",
			},
			expectedScore: 1000,
expectedLevel: "CRITICAL",
			expectedVerd:  "QUARANTINE",
		},
		{
			name: "Score Cap",
			findings: []string{
				"Behavior Rule: PowerShell Downloader Detected",
				"Behavior Rule: Malware Dropper Detected",
				"Behavior Rule: Malicious Office Execution",
				"VirusTotal malicious hash detected",
				"VirusTotal malicious hash detected",
				"VirusTotal malicious hash detected",
				"Persistence Detected: Registry Run Key Modification",
				"Persistence Detected: Registry Run Key Modification",
			},
			expectedScore: 1000,
			expectedLevel: "CRITICAL",
			expectedVerd:  "QUARANTINE",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			score, level, verdict :=
				CalculateSandboxRisk(
					tt.findings,
				)

			if score != tt.expectedScore {
				t.Errorf(
					"expected score %d got %d",
					tt.expectedScore,
					score,
				)
			}

			if level != tt.expectedLevel {
				t.Errorf(
					"expected level %s got %s",
					tt.expectedLevel,
					level,
				)
			}

			if verdict != tt.expectedVerd {
				t.Errorf(
					"expected verdict %s got %s",
					tt.expectedVerd,
					verdict,
				)
			}
		})
	}
}