package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"phishing-platform/database"
)

type NoteRequest struct {
	Note string `json:"note"`
}

func AddNoteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	idParam := strings.TrimPrefix(
		r.URL.Path,
		"/add-note/",
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

	var request NoteRequest

	err = json.NewDecoder(
		r.Body,
	).Decode(&request)

	if err != nil {

		http.Error(
			w,
			"Invalid JSON",
			http.StatusBadRequest,
		)

		return
	}

	query := `
UPDATE public.emails
SET analyst_note = $1
WHERE id = $2
`

	_, err = database.DB.Exec(
		query,
		request.Note,
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

	w.Write([]byte(
		"Note added successfully",
	))
}