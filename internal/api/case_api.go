package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"phishing-platform/database"
)

type CreateCaseRequest struct {
	FileName string `json:"file_name"`
	Analyst  string `json:"analyst"`
}

type UpdateCaseRequest struct {
	Status string `json:"status"`
	Notes  string `json:"notes"`
}

func CreateCaseHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req CreateCaseRequest

	err :=
		json.NewDecoder(
			r.Body,
		).Decode(
			&req,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	err =
		database.CreateCase(
			req.FileName,
			req.Analyst,
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
			"message": "Case created",
		},
	)
}

func GetCasesHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	cases, err :=
		database.GetAllCases()

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

	json.NewEncoder(w).Encode(
		cases,
	)
}

func UpdateCaseHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id :=
		strings.TrimPrefix(
			r.URL.Path,
			"/api/case-details/",
		)

	var req UpdateCaseRequest

	err :=
		json.NewDecoder(
			r.Body,
		).Decode(
			&req,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	err =
		database.UpdateCase(
			id,
			req.Status,
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

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Case updated",
		},
	)
}

func GetCaseHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id :=
		strings.TrimPrefix(
			r.URL.Path,
			"/api/case-details/",
		)

	caseData, err :=
		database.GetCaseByID(
			id,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		caseData,
	)
}

func CloseCaseHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id :=
		strings.TrimPrefix(
			r.URL.Path,
			"/api/case-close/",
		)

	err :=
		database.CloseCase(
			id,
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
			"message": "Case closed",
		},
	)
}