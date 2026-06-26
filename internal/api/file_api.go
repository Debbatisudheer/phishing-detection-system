package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"phishing-platform/internal/macroanalyzer"
"phishing-platform/internal/yara"
"phishing-platform/internal/sandbox"
"strings"
"phishing-platform/internal/risk"
"phishing-platform/internal/qr"
"phishing-platform/internal/zipanalyzer"
"phishing-platform/internal/pdfanalyzer"
"phishing-platform/internal/parser"
"fmt"
"phishing-platform/database"
"phishing-platform/internal/hash"
"phishing-platform/internal/mitre"
"phishing-platform/internal/threatintel"
"phishing-platform/internal/websocket"
"phishing-platform/internal/virustotal"
"phishing-platform/internal/domain"
)

type AnalyzeFileResponse struct {

    FileName      string   `json:"file_name"`
    FilePath      string   `json:"file_path"`
    Findings      []string `json:"findings"`
    RiskScore     int      `json:"risk_score"`
    RiskLevel     string   `json:"risk_level"`
    Verdict       string   `json:"verdict"`
    Message       string   `json:"message"`

    SandboxJobID  int      `json:"sandbox_job_id"`
    SandboxStatus string   `json:"sandbox_status"`

}

func AnalyzeFileHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	file, header, err :=
		r.FormFile(
			"file",
		)

	if err != nil {

		http.Error(
			w,
			"File not provided",
			http.StatusBadRequest,
		)

		return
	}

	defer file.Close()

	os.MkdirAll(
		"uploads",
		0755,
	)

	savePath :=
		filepath.Join(
			"uploads",
			header.Filename,
		)

	dst, err :=
		os.Create(
			savePath,
		)

	if err != nil {

		http.Error(
			w,
			"Failed to save file",
			http.StatusInternalServerError,
		)

		return
	}

	defer dst.Close()

	_, err = io.Copy(
		dst,
		file,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to save file",
			http.StatusInternalServerError,
		)

		return
	}

	var findings []string
	var extractedURLs []string

	// DOCM ANALYSIS
	if strings.HasSuffix(
		strings.ToLower(savePath),
		".docm",
	) {

		macroContent :=
			macroanalyzer.ExtractMacroText(
				savePath,
			)

		macroFindings :=
			macroanalyzer.AnalyzeMacroContent(
				macroContent,
			)

		findings = append(
			findings,
			macroFindings...,
		)

		yaraFindings :=
			yara.ScanContent(
				macroContent,
			)

		findings = append(
			findings,
			yaraFindings...,
		)

		sandboxFindings :=
			sandbox.AnalyzeBehavior(
				macroContent,
			)

		findings = append(
			findings,
			sandboxFindings...,
		)
	}

	if strings.HasSuffix(
    strings.ToLower(savePath),
    ".ps1",
) {

    psFindings :=
        sandbox.AnalyzeSandboxContent(
            savePath,
        )

    findings = append(
        findings,
        psFindings...,
    )
}

	// PNG ANALYSIS
if strings.HasSuffix(
	strings.ToLower(savePath),
	".png",
) {

	qrResults :=
		qr.DecodeQRImage(
			savePath,
		)

	fmt.Println(
		"QR Results:",
		qrResults,
	)

	for _, qrURL := range qrResults {

		err := database.SaveIOC(
	qrURL,
	"QR",
	header.Filename,
)

if err != nil {

	fmt.Println(
		"IOC Save Error:",
		err,
	)
}

	extractedURLs = append(
		extractedURLs,
		qrURL,
	)

	findings = append(
		findings,
		"QR URL extracted: "+qrURL,
	)

	domainFindings :=
		domain.AnalyzeURL(
			qrURL,
		)

	findings = append(
		findings,
		domainFindings...,
	)

	threatFindings :=
		threatintel.CheckThreatIntel(
			qrURL,
		)

	findings = append(
		findings,
		threatFindings...,
	)

	reputationFindings :=
		threatintel.CheckURLReputation(
			qrURL,
		)

	findings = append(
		findings,
		reputationFindings...,
	)
}
}

	// ZIP ANALYSIS
	if strings.HasSuffix(
		strings.ToLower(savePath),
		".zip",
	) {

		files, err :=
			zipanalyzer.ExtractZIPContents(
				savePath,
			)
			_, err =
    zipanalyzer.ExtractArtifactsForSandbox(
        savePath,
    )

if err != nil {

	fmt.Println(
		"Artifact Extraction Error:",
		err,
	)
}

		if err == nil {

			contentFindings :=
				zipanalyzer.AnalyzeZIPFileContents(
					files,
				)

			findings = append(
				findings,
				contentFindings...,
			)

			for _, file := range files {

    lower :=
        strings.ToLower(file)

    if strings.HasSuffix(lower, ".exe") ||
        strings.HasSuffix(lower, ".ps1") ||
        strings.HasSuffix(lower, ".docm") ||
        strings.HasSuffix(lower, ".xlsm") {

        findings = append(
            findings,
            "Sandbox candidate: "+file,
        )
    }
}

			nestedFindings :=
				zipanalyzer.DetectNestedZIP(
					savePath,
				)

			findings = append(
				findings,
				nestedFindings...,
			)
		}
	}

	// PDF ANALYSIS
	if strings.HasSuffix(
		strings.ToLower(savePath),
		".pdf",
	) {

		pdfText :=
	pdfanalyzer.ExtractPDFText(
		savePath,
	)

pdfText = strings.ReplaceAll(
	pdfText,
	"\r\n",
	" ",
)

pdfText = strings.ReplaceAll(
	pdfText,
	"\n",
	" ",
)

pdfText = strings.ReplaceAll(
	pdfText,
	" - ",
	"-",
)

pdfText = strings.ReplaceAll(
	pdfText,
	" -",
	"-",
)

pdfText = strings.ReplaceAll(
	pdfText,
	"- ",
	"-",
)

pdfFindings :=
	pdfanalyzer.AnalyzePDFText(
		pdfText,
	)

		findings = append(
			findings,
			pdfFindings...,
		)

		urls :=
			parser.ExtractURLs(
				pdfText,
			)

			extractedURLs = append(
	extractedURLs,
	urls...,
)

		for _, url := range urls {
			err := database.SaveIOC(
	url,
	"PDF",
	header.Filename,
)

if err != nil {

	fmt.Println(
		"IOC Save Error:",
		err,
	)
}

	findings = append(
		findings,
		"PDF URL extracted: "+url,
	)

	domainFindings :=
	domain.AnalyzeURL(
		url,
	)

findings = append(
	findings,
	domainFindings...,
)

	threatFindings :=
		threatintel.CheckThreatIntel(
			url,
		)

	findings = append(
		findings,
		threatFindings...,
	)

	reputationFindings :=
		threatintel.CheckURLReputation(
			url,
		)

	findings = append(
		findings,
		reputationFindings...,
	)
}
	}
mitreTechniques :=
	mitre.MapFileTechniques(
		findings,
	)
	riskScore :=
		risk.CalculateRisk(
			header.Filename,
			"",
			nil,
			findings,
		)

		

	riskLevel :=
		risk.GetRiskLevel(
			riskScore,
		)

	verdict := "ALLOW"


	if riskLevel == "MEDIUM" {

		verdict = "SUSPICIOUS"

	} else if riskLevel == "HIGH" ||
		riskLevel == "CRITICAL" {

		verdict = "QUARANTINE"
	}

	if riskLevel == "HIGH" ||
	riskLevel == "CRITICAL" {

	alert :=
		fmt.Sprintf(
			"ALERT | File=%s | Risk=%s | Verdict=%s",
			header.Filename,
			riskLevel,
			verdict,
		)

	err :=
		database.SaveAlert(
			header.Filename,
			riskLevel,
			verdict,
			alert,
		)

	if err != nil {

		fmt.Println(
			"Save Alert Error:",
			err,
		)
	}

	websocket.Broadcast <-
		[]byte(alert)
}

	

sha256 :=
	hash.CalculateSHA256(
		savePath,
	)

// Query VirusTotal only if it has NOT already
// been done inside AnalyzeSandboxContent()

if !strings.HasSuffix(
	strings.ToLower(savePath),
	".ps1",
) {

	vtResponse, err :=
		virustotal.QueryHash(
			sha256,
		)

	if err == nil {

		vtFindings :=
			virustotal.CheckHashReputation(
				vtResponse,
			)

		findings = append(
			findings,
			vtFindings...,
		)
	}
}

fmt.Println(
	"SHA256:",
	sha256,
)
uniqueMitres :=
	make(map[string]bool)

for _, technique :=
	range mitreTechniques {

	uniqueMitres[
		technique,
	] = true
}

var cleanMitres []string

for technique :=
	range uniqueMitres {

	cleanMitres = append(
		cleanMitres,
		technique,
	)
}

	err = database.SaveAnalysisResult(
	header.Filename,
	riskScore,
	riskLevel,
	verdict,
	strings.Join(
		findings,
		"\n",
	),
	sha256,
	strings.Join(
		extractedURLs,
		"\n",
	),
	strings.Join(
	cleanMitres,
	", ",
),
)

if err != nil {

	fmt.Println(
		"DB SAVE ERROR:",
		err,
	)
}

lowerFile :=
	strings.ToLower(
		header.Filename,
	)
sandboxJobID := 0
sandboxStatus := "NOT_REQUIRED"

if strings.HasSuffix(
    lowerFile,
    ".exe",
) ||
    strings.HasSuffix(
        lowerFile,
        ".ps1",
    ) ||
    strings.HasSuffix(
        lowerFile,
        ".docm",
    ) ||
    strings.HasSuffix(
        lowerFile,
        ".xlsm",
    ) {

    sandboxStatus = "RUNNING"

    jobID, err :=
        database.CreateSandboxJob(
            header.Filename,
            savePath,
        )

    if err != nil {

        fmt.Println(
            "Sandbox Job Error:",
            err,
        )

    } else {

        sandboxJobID = jobID

    }

}

	response :=
    AnalyzeFileResponse{
        FileName:      header.Filename,
        FilePath:      savePath,
        Findings:      findings,
        RiskScore:     riskScore,
        RiskLevel:     riskLevel,
        Verdict:       verdict,
        Message:       "File analyzed",

        SandboxJobID:  sandboxJobID,
        SandboxStatus: sandboxStatus,
    }

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		response,
	)
}