package sandbox

import (
	"bytes"
	"context"
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

	println(
		"DOCKER FILE:",
		fullPath,
	)

	ctx, cancel :=
		context.WithTimeout(
			context.Background(),
			30*time.Second,
		)

	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"docker",
		"run",

		"--rm",

		"--memory=256m",

		"--cpus=1",

		"-v",
		fullPath+":/sample",

		"ubuntu:22.04",

		"cat",
		"/sample",
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

	if ctx.Err() ==
		context.DeadlineExceeded {

		return out.String(),
			duration,
			ctx.Err()
	}

	return out.String(),
		duration,
		err
}