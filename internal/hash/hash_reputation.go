package hash

var MaliciousHashes = map[string]bool{

	"091b572abf984382a95b9455e407f284e24bfb616890abc3fcdb2be68813f39d": true,

	"22aba92ebdbdd78662e009feacc0f8e3381ae03293a29c16161eed9ef8fb2a3a": true,

	"46f465d64f3d54aeb53f63df6aa1f36dfa50eeac9a787ba5f06ae4531af5f42c": true,
}

func CheckHashReputation(
	hash string,
) []string {

	findings := []string{}

	if MaliciousHashes[hash] {

		findings = append(
			findings,
			"Known malware hash detected",
		)
	}

	return findings
}