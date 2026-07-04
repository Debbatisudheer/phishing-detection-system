package api

import (
	"encoding/json"
	"net/http"

	notesrepo "phishing-platform/database/notes"
)

type AnalystNoteRequest struct {
	IOC     string `json:"ioc"`
	Analyst string `json:"analyst"`
	Notes   string `json:"notes"`
}

func SaveNoteHandler(
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

	var req AnalystNoteRequest

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

	if req.IOC == "" ||
		req.Analyst == "" ||
		req.Notes == "" {

		http.Error(
			w,
			"ioc, analyst and notes are required",
			http.StatusBadRequest,
		)

		return
	}

	err = notesrepo.SaveAnalystNote(
		req.IOC,
		req.Analyst,
		req.Notes,
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
			"message": "Note Saved",
		},
	)
}

func GetNotesHandler(
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

	ioc := r.URL.Query().Get(
		"ioc",
	)

	if ioc == "" {

		http.Error(
			w,
			"ioc is required",
			http.StatusBadRequest,
		)

		return
	}

	data, err := notesrepo.GetAnalystNotes(
		ioc,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	if data == nil {

		data = []map[string]interface{}{}

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
		data,
	)
}