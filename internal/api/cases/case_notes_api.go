package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	caserepo "phishing-platform/database/cases"
)

type NoteRequest struct {
	CaseID  int    `json:"case_id"`
	Analyst string `json:"analyst"`
	Note    string `json:"note"`
}

func AddCaseNoteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Only POST allowed
	if r.Method != http.MethodPost {

		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var req NoteRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(
		&req,
	)

	if err != nil {

		http.Error(
			w,
			"Invalid Request",
			http.StatusBadRequest,
		)

		return
	}

	if req.CaseID == 0 ||
		req.Analyst == "" ||
		req.Note == "" {

		http.Error(
			w,
			"case_id, analyst and note are required",
			http.StatusBadRequest,
		)

		return
	}

	err = caserepo.AddCaseNote(
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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusCreated,
	)

	json.NewEncoder(
		w,
	).Encode(
		map[string]string{
			"message": "Note added",
		},
	)
}

func GetCaseNotesHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Only GET allowed
	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	idStr := strings.TrimPrefix(
		r.URL.Path,
		"/api/case-notes/",
	)

	if idStr == "" {

		http.Error(
			w,
			"Missing case id",
			http.StatusBadRequest,
		)

		return
	}

	caseID, err := strconv.Atoi(
		idStr,
	)

	if err != nil {

		http.Error(
			w,
			"Invalid case id",
			http.StatusBadRequest,
		)

		return
	}

	notes, err := caserepo.GetCaseNotes(
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

	if notes == nil {

		notes = []map[string]interface{}{}

	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(
		w,
	).Encode(
		notes,
	)
}
