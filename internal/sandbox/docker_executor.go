package sandbox

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ExecuteInDocker(
	filePath string,
) (
	string,
	int,
	error,
) {

	absPath, err :=
		os.Getwd()

	if err != nil {
		return "", 0, err
	}

	fullPath :=
		absPath + "\\" + filePath

	rulesPath :=
		absPath + "\\rules"

	fmt.Println(
		"DOCKER FILE:",
		fullPath,
	)

	fmt.Println(
		"RULES PATH:",
		rulesPath,
	)

	ctx, cancel :=
		context.WithTimeout(
			context.Background(),
			120*time.Second,
		)

	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"docker",
		"run",

		"--rm",

		"--memory=512m",

		"--cpus=1",

		"-v",
		fullPath+":/sample",

		"-v",
		rulesPath+":/rules",

		"phishing-sandbox",

		"sh",
		"-c",

		"file /sample ; " +
"strings /sample | head -100 ; " +
"sha256sum /sample ; " +
"timeout 45 clamscan --no-summary /sample ; " +
"yara /rules/malware.yar /sample",
	)

	var out bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &out

	start :=
		time.Now()

	err = cmd.Run()

	duration :=
		int(
			time.Since(
				start,
			).Seconds(),
		)

	fmt.Println(
		"DOCKER EXIT ERROR:",
		err,
	)

	fmt.Println(
		"DOCKER OUTPUT LENGTH:",
		len(out.String()),
	)

	if ctx.Err() != nil {

		fmt.Println(
			"CONTEXT ERROR:",
			ctx.Err(),
		)
	}

	fmt.Println(
		"COMMAND:",
		cmd.String(),
	)

	fmt.Println(
		"========== RAW DOCKER RESULT ==========",
	)

	fmt.Println(
		out.String(),
	)

	fmt.Println(
		"=======================================",
	)

	if ctx.Err() ==
		context.DeadlineExceeded {

		return out.String(),
			duration,
			ctx.Err()
	}

	// YARA often returns exit status 1 when matches occur.
	// We still want the output.
	return out.String(),
		duration,
		nil
}