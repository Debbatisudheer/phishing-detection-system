package sandbox

import (
	"fmt"
	"strings"
	"time"

	"phishing-platform/database"
	"phishing-platform/internal/hash"
)

func StartSandboxWorker() {

	go func() {

		for {

			jobs, err :=
				database.GetSandboxJobs()

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
					database.UpdateSandboxJobStatus(
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
					database.SaveSandboxReport(
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
				}

				err =
					database.UpdateSandboxJobStatus(
						job.ID,
						"COMPLETED",
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