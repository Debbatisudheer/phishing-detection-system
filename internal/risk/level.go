package risk

func GetRiskLevel(
	score int,
) string {

	switch {

	case score >= 500:
		return "CRITICAL"

	case score >= 300:
		return "HIGH"

	case score >= 100:
		return "MEDIUM"

	default:
		return "LOW"
	}
}