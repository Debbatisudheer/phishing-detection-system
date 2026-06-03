package hash

var MaliciousHashes = map[string]bool{

	"091b572abf984382a95b9455e407f284e24bfb616890abc3fcdb2be68813f39d": true,
}

func CheckHashReputation(
	hash string,
) []string {

	var findings []string

	if MaliciousHashes[hash] {

		findings = append(
			findings,
			"Known malware hash detected",
		)
	}

	return findings
}