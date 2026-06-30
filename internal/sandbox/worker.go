package sandbox

import (
	"fmt"
	"strings"
	"time"
	sandboxrepo "phishing-platform/database/sandbox"
	"phishing-platform/internal/hash"
)

func StartSandboxWorker() {

	go func() {

		for {

			jobs, err :=
				sandboxrepo.GetSandboxJobs()

			if err != nil {

				fmt.Println(
					"Sandbox Worker Error:",
					err,
				)

				time.Sleep(
					10 * time.Second,
				)

				continue
			}

			for _, job := range jobs {

				if job.Status != "PENDING" {
					continue
				}

				fmt.Println(
					"Processing Sandbox Job:",
					job.ID,
				)

				err :=
					sandboxrepo.UpdateSandboxJobStatus(
						job.ID,
						"RUNNING",
					)

				if err != nil {

					fmt.Println(
						"Status Update Error:",
						err,
					)

					continue
				}

				time.Sleep(
					5 * time.Second,
				)
				contentFindings :=
	AnalyzeSandboxContent(
		job.FilePath,
	)

fmt.Println(
	"SANDBOX PATH:",
		job.FilePath,
)

output,
duration,
err :=
	ExecuteInDocker(
		job.FilePath,
	)

	fmt.Println(
    "========== DOCKER OUTPUT ==========",
)
fmt.Println(
    output,
)
fmt.Println(
    "===================================",
)

fmt.Println(
	"DOCKER OUTPUT:",
	output,
)

dockerFindings :=
	BuildDockerReport(
		output,
		err,
	)

	analysisFindings :=
	AnalyzeDockerOutput(
		output,
	)

contentFindings = append(
	contentFindings,
	analysisFindings...,
)

	executionStatus := "SUCCESS"

if err != nil {

	executionStatus = "FAILED"
}

err = sandboxrepo.SaveDockerReport(
	job.ID,
	"DESTROYED",
	executionStatus,
	duration,
)

if err != nil {

	fmt.Println(
		"Docker Report Error:",
		err,
	)
}

contentFindings = append(
	contentFindings,
	dockerFindings...,
)
contentFindings = append(
	contentFindings,
	"Docker Duration: "+
		fmt.Sprint(duration)+
		" seconds",
)

timeline :=
	BuildTimeline(
		contentFindings,
	)

fmt.Println(
	"Sandbox Timeline:",
)

for _, event := range timeline {

	fmt.Println(
		event,
	)
}
					metadata :=
	GetFileMetadata(
		job.FilePath,
	)
	sha256 :=
	hash.CalculateSHA256(
		job.FilePath,
	)

fmt.Println(
	"File Name:",
	metadata.FileName,
)

fmt.Println(
	"File Size:",
	metadata.FileSize,
)

fmt.Println(
	"Extension:",
	metadata.Extension,
)

fmt.Println(
	"MIME Type:",
	metadata.MIMEType,
)

fmt.Println(
	"MD5:",
	metadata.MD5,
)

				fmt.Println(
					"Sandbox Findings:",
					contentFindings,
				)

				riskScore,
					riskLevel,
					verdict :=
					CalculateSandboxRisk(
						contentFindings,
					)

					fmt.Println(
    "RISK SCORE:",
    riskScore,
)

fmt.Println(
    "RISK LEVEL:",
    riskLevel,
)

fmt.Println(
    "VERDICT:",
    verdict,
)

				mitre :=
					MapSandboxMITRE(
						contentFindings,
					)

				findings :=
					strings.Join(
						contentFindings,
						"\n",
					)

					fmt.Println(
	"FINAL FINDINGS STRING:",
	findings,
)

				err =
    sandboxrepo.SaveSandboxReport(
        job.ID,
        metadata.FileName,
        metadata.FileSize,
        metadata.Extension,
        metadata.MIMEType,
        metadata.MD5,
        sha256,
        findings,
        riskScore,
        riskLevel,
        verdict,
        mitre,
    )

if err != nil {

    fmt.Println(
        "Sandbox Report Error:",
        err,
    )

} else {

    fmt.Println(
        "SANDBOX REPORT SAVED SUCCESSFULLY",
    )

    fmt.Println(
        "JOB ID:",
        job.ID,
    )

    fmt.Println(
        "RISK SCORE:",
        riskScore,
    )

    fmt.Println(
        "RISK LEVEL:",
        riskLevel,
    )
}

				status := "COMPLETED"

if riskLevel == "HIGH" ||
    riskLevel == "CRITICAL" {

    status = "MALICIOUS"

} else if riskLevel == "MEDIUM" {

    status = "SUSPICIOUS"
}

fmt.Println(
    "SANDBOX RISK SCORE:",
    riskScore,
)

fmt.Println(
    "SANDBOX RISK LEVEL:",
    riskLevel,
)

fmt.Println(
    "SANDBOX STATUS:",
    status,
)

err =
    sandboxrepo.UpdateSandboxJobStatus(
        job.ID,
        status,
    )

				if err != nil {

					fmt.Println(
						"Completion Update Error:",
						err,
					)
				}

				fmt.Println(
					"Sandbox Job Completed:",
					job.ID,
				)
			}

			time.Sleep(
				10 * time.Second,
			)
		}
	}()
}