package decision

func MakeDecision(riskScore int) string {

	if riskScore >= 70 {
		return "QUARANTINE"
	}

	if riskScore >= 40 {
		return "WARNING"
	}

	return "ALLOW"
}