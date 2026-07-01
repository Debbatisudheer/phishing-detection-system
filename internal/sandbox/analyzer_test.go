package sandbox

import "testing"

func TestAnalyzeSandboxFile(t *testing.T) {

	tests := []struct {
		name       string
		fileName   string
		expected   SandboxAnalysis
	}{
		{
			name:     "Executable",
			fileName: "malware.exe",
			expected: SandboxAnalysis{
				Findings:  "Executable Detected",
				RiskScore: 300,
				RiskLevel: "HIGH",
				Verdict:   "QUARANTINE",
				Mitre:     "T1204 - User Execution",
			},
		},
		{
			name:     "PowerShell",
			fileName: "payload.ps1",
			expected: SandboxAnalysis{
				Findings:  "PowerShell Script Detected",
				RiskScore: 500,
				RiskLevel: "CRITICAL",
				Verdict:   "QUARANTINE",
				Mitre:     "T1059.001 - PowerShell",
			},
		},
		{
			name:     "DOCM",
			fileName: "invoice.docm",
			expected: SandboxAnalysis{
				Findings:  "Macro Document Detected",
				RiskScore: 350,
				RiskLevel: "HIGH",
				Verdict:   "QUARANTINE",
				Mitre:     "T1566.001 - Spearphishing Attachment",
			},
		},
		{
			name:     "XLSM",
			fileName: "sheet.xlsm",
			expected: SandboxAnalysis{
				Findings:  "Macro Spreadsheet Detected",
				RiskScore: 350,
				RiskLevel: "HIGH",
				Verdict:   "QUARANTINE",
				Mitre:     "T1566.001 - Spearphishing Attachment",
			},
		},
		{
			name:     "Unknown File",
			fileName: "photo.jpg",
			expected: SandboxAnalysis{},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := AnalyzeSandboxFile(tt.fileName)

			if result.Findings != tt.expected.Findings {
				t.Errorf("Findings: expected %q got %q",
					tt.expected.Findings,
					result.Findings)
			}

			if result.RiskScore != tt.expected.RiskScore {
				t.Errorf("RiskScore: expected %d got %d",
					tt.expected.RiskScore,
					result.RiskScore)
			}

			if result.RiskLevel != tt.expected.RiskLevel {
				t.Errorf("RiskLevel: expected %q got %q",
					tt.expected.RiskLevel,
					result.RiskLevel)
			}

			if result.Verdict != tt.expected.Verdict {
				t.Errorf("Verdict: expected %q got %q",
					tt.expected.Verdict,
					result.Verdict)
			}

			if result.Mitre != tt.expected.Mitre {
				t.Errorf("MITRE: expected %q got %q",
					tt.expected.Mitre,
					result.Mitre)
			}
		})
	}
}