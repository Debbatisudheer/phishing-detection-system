package timeline

import (
	"fmt"
	"time"
)

func LogEvent(
	event string,
) {

	fmt.Printf(
		"[%s] %s\n",
		time.Now().Format(
			"2006-01-02 15:04:05",
		),
		event,
	)
}