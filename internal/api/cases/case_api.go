package api

import (
	"encoding/json"
	"net/http"
	"strings"

	caserepo "phishing-platform/database/cases"
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

	// Only POST allowed
	if r.Method != http.MethodPost {

		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var req CreateCaseRequest

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

	if req.FileName == "" ||
		req.Analyst == "" {

		http.Error(
			w,
			"file_name and analyst are required",
			http.StatusBadRequest,
		)

		return
	}

	err = caserepo.CreateCase(
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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusCreated,
	)

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

	// Only GET allowed
	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	cases, err :=
		caserepo.GetAllCases()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	if cases == nil {

		cases = []map[string]interface{}{}

	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(w).Encode(
		cases,
	)
}

func UpdateCaseHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Only PUT allowed
	if r.Method != http.MethodPut {

		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	// Fixed route
	id := strings.TrimPrefix(
		r.URL.Path,
		"/api/case/",
	)

	if id == "" {

		http.Error(
			w,
			"Missing case id",
			http.StatusBadRequest,
		)

		return
	}

	var req UpdateCaseRequest

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

	err = caserepo.UpdateCase(
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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

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

	// Only GET allowed
	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	id := strings.TrimPrefix(
		r.URL.Path,
		"/api/case-details/",
	)

	if id == "" {

		http.Error(
			w,
			"Missing case id",
			http.StatusBadRequest,
		)

		return
	}

	caseData, err :=
		caserepo.GetCaseByID(
			id,
		)

	if err != nil {

		http.Error(
			w,
			"Case not found",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(w).Encode(
		caseData,
	)
}

func CloseCaseHandler(
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

	id := strings.TrimPrefix(
		r.URL.Path,
		"/api/case-close/",
	)

	if id == "" {

		http.Error(
			w,
			"Missing case id",
			http.StatusBadRequest,
		)

		return
	}

	err := caserepo.CloseCase(
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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Case closed",
		},
	)
}
