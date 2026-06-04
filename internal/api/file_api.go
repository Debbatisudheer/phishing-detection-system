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
)

type AnalyzeFileResponse struct {
	FileName  string   `json:"file_name"`
	FilePath  string   `json:"file_path"`
	Findings  []string `json:"findings"`
	RiskScore int      `json:"risk_score"`
	RiskLevel string   `json:"risk_level"`
	Verdict   string   `json:"verdict"`
	Message   string   `json:"message"`
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

	// PNG ANALYSIS
	if strings.HasSuffix(
		strings.ToLower(savePath),
		".png",
	) {

		qrResults :=
			qr.DecodeQRImage(
				savePath,
			)

		for _, qrURL := range qrResults {

	extractedURLs = append(
		extractedURLs,
		qrURL,
	)

	findings = append(
		findings,
		"QR URL extracted: "+qrURL,
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

		if err == nil {

			contentFindings :=
				zipanalyzer.AnalyzeZIPFileContents(
					files,
				)

			findings = append(
				findings,
				contentFindings...,
			)

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

	findings = append(
		findings,
		"PDF URL extracted: "+url,
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

	sha256 :=
	hash.CalculateSHA256(
		savePath,
	)

	fmt.Println(
	"SHA256:",
	sha256,
)


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
		mitreTechniques,
		"\n",
	),
)

if err != nil {

	fmt.Println(
		"DB SAVE ERROR:",
		err,
	)
}
	response :=
		AnalyzeFileResponse{
			FileName:  header.Filename,
			FilePath:  savePath,
			Findings:  findings,
			RiskScore: riskScore,
			RiskLevel: riskLevel,
			Verdict:   verdict,
			Message:   "File analyzed",
		}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		response,
	)
}