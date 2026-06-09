package decision

func MakeDecision(
	riskScore int,
) string {

	if riskScore >= 400 {

		return "QUARANTINE"
	}

	if riskScore >= 100 {

		return "SUSPICIOUS"
	}

	return "ALLOW"
}