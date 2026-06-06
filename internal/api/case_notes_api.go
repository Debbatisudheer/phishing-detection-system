package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"phishing-platform/database"
)

type NoteRequest struct {
	CaseID int    `json:"case_id"`
	Analyst string `json:"analyst"`
	Note string `json:"note"`
}

func AddCaseNoteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req NoteRequest

	json.NewDecoder(
		r.Body,
	).Decode(
		&req,
	)

	err :=
		database.AddCaseNote(
			req.CaseID,
			req.Analyst,
			req.Note,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message":
				"Note added",
		},
	)
}

func GetCaseNotesHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	idStr :=
		strings.TrimPrefix(
			r.URL.Path,
			"/api/case-notes/",
		)

	caseID, _ :=
		strconv.Atoi(
			idStr,
		)

	notes, err :=
		database.GetCaseNotes(
			caseID,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		notes,
	)
}