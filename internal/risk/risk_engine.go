package risk

import (
	"fmt"
	"strings"
)

func CalculateRisk(
	subject string,
	body string,
	urls []string,
	findings []string,
) int {

	score := 0

	// URL existence
	if len(urls) > 0 {
		score += 40
	}

	// Analyze findings
	for _, finding := range findings {

		fmt.Println(
			"SCORING FINDING:",
			finding,
		)

		for pattern, weight := range RiskWeights {

			if strings.Contains(
				strings.ToLower(finding),
				strings.ToLower(pattern),
			) {
				score += weight
			}
		}

		fmt.Println(
			"CURRENT SCORE:",
			score,
		)
	}

	// Normalize text
	subject = strings.ToLower(subject)
	body = strings.ToLower(body)

	// Subject keywords
	for keyword, weight := range SubjectWeights {

		if strings.Contains(
			subject,
			keyword,
		) {
			score += weight
		}
	}

	// Body keywords
	for keyword, weight := range BodyWeights {

		if strings.Contains(
			body,
			keyword,
		) {
			score += weight
		}
	}

	// Maximum score cap
	if score > 1000 {
		score = 1000
	}

	fmt.Println(
		"FINAL RISK SCORE:",
		score,
	)

	fmt.Println(
		"RISK LEVEL:",
		GetRiskLevel(score),
	)

	return score
}