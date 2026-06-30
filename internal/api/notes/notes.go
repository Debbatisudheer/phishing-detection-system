package api

import (
	"encoding/json"
	"net/http"
	notesrepo "phishing-platform/database/notes"
)

type AnalystNoteRequest struct {
	IOC string `json:"ioc"`
	Analyst string `json:"analyst"`
	Notes string `json:"notes"`
}

func SaveNoteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

var req AnalystNoteRequest

	json.NewDecoder(
		r.Body,
	).Decode(
		&req,
	)

	err :=
		notesrepo.SaveAnalystNote(
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

	w.Write(
		[]byte(
			"Note Saved",
		),
	)
}

func GetNotesHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	ioc :=
		r.URL.Query().Get(
			"ioc",
		)

	data, err :=
		notesrepo.GetAnalystNotes(
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

	json.NewEncoder(w).Encode(
		data,
	)
}