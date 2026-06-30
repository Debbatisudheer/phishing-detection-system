package auth

import (
    "encoding/json"
    "net/http"

    userrepo "phishing-platform/database/users"
    "phishing-platform/internal/jwt"
)

func RegisterHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req RegisterRequest

	err := json.NewDecoder(
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

	err = userrepo.CreateUser(
		req.Username,
		req.Password,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	token, err :=
		jwt.GenerateToken(
			req.Username,
			"ANALYST",
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
			"token": token,
		},
	)
}
