package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"phishing-platform/database"
)

func QuarantineHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	idParam := strings.TrimPrefix(
		r.URL.Path,
		"/quarantine/",
	)

	id, err := strconv.Atoi(idParam)

	if err != nil {

		http.Error(
			w,
			"Invalid ID",
			http.StatusBadRequest,
		)

		return
	}

	query := `
UPDATE public.emails
SET decision = 'QUARANTINE'
WHERE id = $1
`

	_, err = database.DB.Exec(
		query,
		id,
	)

	if err != nil {

		http.Error(
			w,
			"Database Error",
			http.StatusInternalServerError,
		)

		return
	}

	fmt.Println(
		"Email quarantined:",
		id,
	)

	w.Write([]byte(
		"Email quarantined successfully",
	))
}