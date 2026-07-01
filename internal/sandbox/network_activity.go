package sandbox

import "strings"

func AnalyzeNetworkActivity(
    content string,
) []string {

    var findings []string

    content = strings.ToLower(content)

    if strings.Contains(content, "http://") ||
        strings.Contains(content, "https://") {

        findings = append(
            findings,
            "Network Activity: Outbound Connection Detected",
        )

        findings = append(
            findings,
            "Network Activity: Internet Communication",
        )
    }

    if strings.Contains(
        content,
        "invoke-webrequest",
    ) {

        findings = append(
            findings,
            "Network Activity: Potential Payload Download",
        )
    }

    return findings
}