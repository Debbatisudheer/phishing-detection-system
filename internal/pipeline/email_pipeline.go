package pipeline

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"phishing-platform/database"
	iocrepo "phishing-platform/database/ioc"
	"phishing-platform/internal/attachment"
	"phishing-platform/internal/decision"
	"phishing-platform/internal/domain"
	"phishing-platform/internal/mitre"
	"phishing-platform/internal/parser"
	"phishing-platform/internal/risk"
	"phishing-platform/internal/threatintel"
	"phishing-platform/internal/websocket"
	"phishing-platform/internal/emailauth"
	"phishing-platform/internal/sender"
	"phishing-platform/internal/campaign"
	"phishing-platform/internal/qr"
	"phishing-platform/internal/hash"
	"phishing-platform/internal/zipanalyzer"
	"phishing-platform/internal/timeline"
	"phishing-platform/internal/splunk"
	"phishing-platform/internal/sigma"
	"phishing-platform/internal/report"
	"phishing-platform/internal/virustotal"
	"phishing-platform/internal/stix"
	"phishing-platform/internal/header"
	"phishing-platform/internal/bec"
	"phishing-platform/internal/pdfanalyzer"
	"phishing-platform/internal/macroanalyzer"
	"phishing-platform/internal/yara"
	"phishing-platform/internal/sandbox"
	"phishing-platform/internal/ioc"
	"phishing-platform/internal/urlanalyzer"
	"phishing-platform/internal/thread"
	"phishing-platform/internal/threatfeed"
	"phishing-platform/internal/detonation"
	"phishing-platform/internal/phishtank"
)

func ProcessEmail(
	senderEmail string,
	replyTo string,
	returnPath string,
	subject string,
	body string,
	attachments []string,
) {

	fmt.Println(
		"=== Processing Email ===",
	)

	timeline.LogEvent(
	"Email Received",
)

	// URL Extraction
	urls := parser.ExtractURLs(body)
	for _, url := range urls {

	err := iocrepo.SaveIOC(
		url,
		"EMAIL",
		subject,
	)

	if err != nil {

		fmt.Println(
			"IOC SAVE ERROR:",
			err,
		)
	}

	fmt.Println(
		"SAVING IOC:",
		url,
	)
}

	fmt.Println("URLs:", urls)

	var allFindings []string

	yaraBodyFindings :=
	yara.ScanContent(
		body,
	)
	sandboxBodyFindings :=
	sandbox.AnalyzeBehavior(
		body,
	)

fmt.Println(
	"Sandbox Body Findings:",
	sandboxBodyFindings,
)

allFindings = append(
	allFindings,
	sandboxBodyFindings...,
)

fmt.Println(
	"YARA Body Findings:",
	yaraBodyFindings,
)

allFindings = append(
	allFindings,
	yaraBodyFindings...,
)

	

	headerFindings :=
	header.AnalyzeHeaders(
		senderEmail,
		replyTo,
		returnPath,
	)

fmt.Println(
	"Header Findings:",
	headerFindings,
)

allFindings = append(
	allFindings,
	headerFindings...,
)
displayNameFindings :=
	header.DetectDisplayNameSpoofing(
		senderEmail,
	)

fmt.Println(
	"Display Name Findings:",
	displayNameFindings,
)

allFindings = append(
	allFindings,
	displayNameFindings...,
)

becFindings :=
	bec.DetectBEC(
		subject,
		body,
	)

fmt.Println(
	"BEC Findings:",
	becFindings,
)

threadFindings :=
	thread.DetectThreadHijack(
		subject,
		body,
	)

fmt.Println(
	"Thread Findings:",
	threadFindings,
)

allFindings = append(
	allFindings,
	threadFindings...,
)

allFindings = append(
	allFindings,
	becFindings...,
)
	// Email Authentication
	authFindings :=
		emailauth.CheckEmailAuthentication(
			senderEmail,
		)

	allFindings = append(
		allFindings,
		authFindings...,
	)

	fmt.Println(
		"Email Authentication:",
		authFindings,
	)

	for _, finding := range authFindings {

	timeline.LogEvent(
		finding,
	)
}

	// Sender Reputation
	senderFindings :=
		sender.CheckSenderReputation(
			senderEmail,
		)

	fmt.Println(
		"Sender Findings:",
		senderFindings,
	)

	for _, finding := range senderFindings {

	timeline.LogEvent(
		finding,
	)
}

	allFindings = append(
		allFindings,
		senderFindings...,
	)

	// Sender History
historyFindings :=
	sender.CheckSenderHistory(
		senderEmail,
	)

fmt.Println(
	"Sender History Findings:",
	historyFindings,
)

for _, finding := range historyFindings {
	timeline.LogEvent(finding)
}
allFindings = append(
	allFindings,
	historyFindings...,
)

campaignFindings :=
	campaign.DetectCampaign(
		subject,
	)

fmt.Println(
	"Campaign Findings:",
	campaignFindings,
)

for _, finding := range campaignFindings {

	timeline.LogEvent(
		finding,
	)
}

allFindings = append(
	allFindings,
	campaignFindings...,
)

qrFindings :=
	qr.DetectQRPhishing(
		subject,
		body,
	)

fmt.Println(
	"QR Findings:",
	qrFindings,
)
for _, finding := range qrFindings {

	timeline.LogEvent(
		finding,
	)
}
allFindings = append(
	allFindings,
	qrFindings...,
)

passwordZipFindings :=
	zipanalyzer.DetectPasswordProtectedZIP(
		subject,
		body,
	)

fmt.Println(
	"Password ZIP Findings:",
	passwordZipFindings,
)

allFindings = append(
	allFindings,
	passwordZipFindings...,
)
	// Domain + Threat Intel
	for _, extractedURL := range urls {

		findings :=
			domain.AnalyzeURL(
				extractedURL,
			)

		fmt.Println(
			"Domain Findings:",
			findings,
		)

		shortenerFindings :=
	urlanalyzer.DetectShortenedURL(
		extractedURL,
	)

fmt.Println(
	"Shortener Findings:",
	shortenerFindings,
)

findings = append(
	findings,
	shortenerFindings...,
)

		

		threatIntelFindings :=
			threatintel.CheckThreatIntel(
				extractedURL,
			)

		urlReputationFindings :=
	threatintel.CheckURLReputation(
		extractedURL,
	)

fmt.Println(
	"URL Reputation Findings:",
	urlReputationFindings,
)

vtResponse, vtErr :=
	virustotal.QueryURL(
		extractedURL,
	)

if vtErr == nil {

	vtFindings :=
		virustotal.CheckURLReputation(
			vtResponse,
		)

	fmt.Println(
		"VirusTotal URL Findings:",
		vtFindings,
	)

	allFindings = append(
		allFindings,
		vtFindings...,
	)

} else {

	fmt.Println(
		"VirusTotal URL Error:",
		vtErr,
	)
}

findings = append(
	findings,
	urlReputationFindings...,
)

		fmt.Println(
			"Threat Intel Findings:",
			threatIntelFindings,
		)

		feedFindings :=
	threatfeed.CheckThreatFeed(
		extractedURL,
	)
phishTankFindings :=
	phishtank.CheckPhishTank(
		extractedURL,
	)

fmt.Println(
	"PhishTank Findings:",
	phishTankFindings,
)

allFindings = append(
	allFindings,
	phishTankFindings...,
)
fmt.Println(
	"Threat Feed Findings:",
	feedFindings,
)

findings = append(
	findings,
	feedFindings...,
)

		findings = append(
			findings,
			threatIntelFindings...,
		)

		allFindings = append(
			allFindings,
			findings...,
		)
	}

	// UEBA
	var uebaFindings []string

		fmt.Println(
	"UEBA Findings:",
	uebaFindings,
)

for _, finding := range uebaFindings {
	timeline.LogEvent(finding)
}

	allFindings = append(
		allFindings,
		uebaFindings...,
	)

	// Attachment Analysis
	// Attachment Analysis
attachmentFindings :=
	attachment.AnalyzeAttachments(
		attachments,
	)

	for _, attachmentFile := range attachments {

	if strings.HasSuffix(
		strings.ToLower(
			attachmentFile,
		),
		".zip",
	) {

		zipFiles, err :=
			zipanalyzer.ExtractZIPContents(
				attachmentFile,
			)

		nestedFindings :=
	zipanalyzer.DetectNestedZIP(
		attachmentFile,
	)

fmt.Println(
	"Nested ZIP Findings:",
	nestedFindings,
)

allFindings = append(
	allFindings,
	nestedFindings...,
)
		
		if err != nil {

			fmt.Println(
				"ZIP Extraction Error:",
				err,
			)

			continue
		}

		fmt.Println(
			"ZIP CONTENTS:",
			zipFiles,
		)

		zipAttachmentFindings :=
			attachment.AnalyzeAttachments(
				zipFiles,
			)

		fmt.Println(
			"ZIP Attachment Findings:",
			zipAttachmentFindings,
		)

		allFindings = append(
			allFindings,
			zipAttachmentFindings...,
		)
	}
}

	fmt.Println(
	"Attachment Findings:",
	attachmentFindings,
)

for _, attachmentFile := range attachments {

	detonationFindings :=
		detonation.AnalyzeAttachmentBehavior(
			attachmentFile,
		)

	fmt.Println(
		"Detonation Findings:",
		detonationFindings,
	)

	allFindings = append(
		allFindings,
		detonationFindings...,
	)
}


allFindings = append(
	allFindings,
	attachmentFindings...,
)

for _, attachmentFile := range attachments {

	if strings.HasSuffix(
		strings.ToLower(attachmentFile),
		".pdf",
	) {

		pdfText :=
			pdfanalyzer.ExtractPDFText(
				attachmentFile,
			)

		fmt.Println(
			"EXTRACTED PDF TEXT:",
			pdfText,
		)

		pdfFindings :=
			pdfanalyzer.AnalyzePDFText(
				pdfText,
			)

		fmt.Println(
			"PDF Findings:",
			pdfFindings,
		)

		allFindings = append(
			allFindings,
			pdfFindings...,
		)

		yaraPDFFindings :=
			yara.ScanContent(
				pdfText,
			)

		fmt.Println(
			"YARA PDF Findings:",
			yaraPDFFindings,
		)

		allFindings = append(
			allFindings,
			yaraPDFFindings...,
		)

		sandboxPDFFindings :=
			sandbox.AnalyzeBehavior(
				pdfText,
			)

		fmt.Println(
			"Sandbox PDF Findings:",
			sandboxPDFFindings,
		)

		allFindings = append(
			allFindings,
			sandboxPDFFindings...,
		)
	}
}
for _, attachmentFile := range attachments {

	if strings.HasSuffix(
		strings.ToLower(attachmentFile),
		".docm",
	) ||
		strings.HasSuffix(
			strings.ToLower(attachmentFile),
			".xlsm",
		) {

		macroContent :=
	macroanalyzer.ExtractWPSMacroText(
		attachmentFile,
	)
	fmt.Println(
	"EXTRACTED MACRO CONTENT:",
	macroContent,
)

		macroFindings :=
			macroanalyzer.AnalyzeMacroContent(
				macroContent,
			)

		fmt.Println(
			"Macro Findings:",
			macroFindings,
		)

		allFindings = append(
			allFindings,
			macroFindings...,
		)

		yaraMacroFindings :=
			yara.ScanContent(
				macroContent,
			)

			sandboxMacroFindings :=
	sandbox.AnalyzeBehavior(
		macroContent,
	)

fmt.Println(
	"Sandbox Macro Findings:",
	sandboxMacroFindings,
)

allFindings = append(
	allFindings,
	sandboxMacroFindings...,
)

		fmt.Println(
			"YARA Macro Findings:",
			yaraMacroFindings,
		)

		allFindings = append(
			allFindings,
			yaraMacroFindings...,
		)
	}
}
for _, attachmentFile := range attachments {

	if strings.HasSuffix(
		strings.ToLower(attachmentFile),
		".png",
	) ||
		strings.HasSuffix(
			strings.ToLower(attachmentFile),
			".jpg",
		) ||
		strings.HasSuffix(
			strings.ToLower(attachmentFile),
			".jpeg",
		) {

		qrResults :=
			qr.DecodeQRImage(
				attachmentFile,
			)

		fmt.Println(
			"QR Image Results:",
			qrResults,
		)

		for _, qrURL := range qrResults {

			allFindings = append(
				allFindings,
				"QR URL extracted: "+qrURL,
			)

			urlFindings :=
				domain.AnalyzeURL(
					qrURL,
				)

			allFindings = append(
				allFindings,
				urlFindings...,
			)
		}
	}
}
	for _, attachmentFile := range attachments {

	fileHash :=
		hash.CalculateSHA256(
			attachmentFile,
		)

	fmt.Println(
		"SHA256:",
		attachmentFile,
		"=>",
		fileHash,
	)

	vtResponse, vtErr :=
	virustotal.QueryHash(
		fileHash,
	)

if vtErr == nil {

	vtFindings :=
		virustotal.CheckHashReputation(
			vtResponse,
		)

	fmt.Println(
		"VirusTotal Findings:",
		vtFindings,
	)

	allFindings = append(
		allFindings,
		vtFindings...,
	)

} else {

	fmt.Println(
		"VirusTotal Error:",
		vtErr,
	)
}

	hashFindings :=
		hash.CheckHashReputation(
			fileHash,
		)

	fmt.Println(
		"Hash Findings:",
		hashFindings,
	)

	allFindings = append(
	allFindings,
	hashFindings...,
)
	}

	// Risk Scoring
	riskScore := risk.CalculateRisk(
		subject,
		body,
		urls,
		allFindings,
	)

	fmt.Println(
		"Risk Score:",
		riskScore,
	)

	timeline.LogEvent(
	fmt.Sprintf(
		"Risk Score = %d",
		riskScore,
	),
)

	fmt.Println(
	"Risk Level:",
	risk.GetRiskLevel(
		riskScore,
	),
)



	// Decision
decisionResult :=
	decision.MakeDecision(
		riskScore,
	)

fmt.Println(
	"Decision:",
	decisionResult,
)

timeline.LogEvent(
	"Decision = " +
	decisionResult,
)

// MITRE Mapping
mitreTechnique :=
	mitre.MapTechnique(
		subject,
		body,
	)

fmt.Println(
	"MITRE Technique:",
	mitreTechnique,
)

riskLevel :=
	risk.GetRiskLevel(
		riskScore,
	)

iocReport :=
	ioc.IOCReport{
		Sender:      senderEmail,
		URLs:        urls,
		Domains:     []string{},
		Hashes:      []string{},
		Attachments: attachments,
		MITRE:       mitreTechnique,
		RiskScore:   riskScore,
		RiskLevel:   riskLevel,
	}

iocErr :=
	ioc.ExportIOC(
		iocReport,
		"ioc_report.json",
	)

if iocErr != nil {

	fmt.Println(
		"IOC Export Error:",
		iocErr,
	)

} else {

	fmt.Println(
		"IOC Report Exported",
	)
}
if len(urls) > 0 {

	stixErr :=
		stix.ExportURLIndicator(
			urls[0],
			"stix_indicator.json",
		)

	if stixErr != nil {

		fmt.Println(
			"STIX Export Error:",
			stixErr,
		)

	} else {

		fmt.Println(
			"STIX Indicator Generated",
		)
	}
}

	reportErr := report.GenerateReport(
	senderEmail,
	subject,
	urls,
	allFindings,
	riskScore,
	risk.GetRiskLevel(
		riskScore,
	),
	decisionResult,
	mitreTechnique,
)
if reportErr != nil {

	fmt.Println(
		"Report Generation Error:",
		reportErr,
	)

} else {

	fmt.Println(
		"Investigation Report Generated",
	)
}

	sigmaErr := sigma.GenerateRule(
	senderEmail,
	riskScore,
	"attack.t1566",
	"phishing_rule.yml",
)

if sigmaErr != nil {

	fmt.Println(
		"Sigma Generation Error:",
		sigmaErr,
	)

} else {

	fmt.Println(
		"Sigma Rule Generated",
	)
}
splunkEvent := splunk.SplunkEvent{
	Timestamp: time.Now().Format(
		"2006-01-02 15:04:05",
	),
	EventType: "phishing_detected",
	Sender:    senderEmail,
	Subject:   subject,
	RiskScore: riskScore,
	Decision:  decisionResult,
	MITRE:     mitreTechnique,
}

exportErr := splunk.ExportEvent(
	splunkEvent,
)

if exportErr != nil {

	fmt.Println(
		"Splunk Export Error:",
		exportErr,
	)

} else {

	fmt.Println(
		"Splunk Event Exported",
	)
}
	// Save to DB
	query := `
INSERT INTO public.emails (
	sender,
	subject,
	body,
	risk_score,
	decision,
	findings,
	attachments,
	mitre_technique
)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8
)
`

	_, err := database.DB.Exec(
	query,
	senderEmail,
	subject,
	body,
	riskScore,
	decisionResult,
	strings.Join(
		allFindings,
		", ",
	),
	strings.Join(
		attachments,
		", ",
	),
	mitreTechnique,
)

	if err != nil {

		fmt.Println(
			"Database Insert Error:",
			err,
		)

		return
	}

	fmt.Println(
		"Email stored successfully",
	)

	// WebSocket Broadcast
	event := map[string]interface{}{
		"sender":          senderEmail,
		"subject":         subject,
		"risk_score":      riskScore,
		"decision":        decisionResult,
		"mitre_technique": mitreTechnique,
	}

	eventJSON, _ :=
		json.Marshal(event)

	websocket.Broadcast <-
		eventJSON
}