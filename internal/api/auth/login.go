package auth

import (
    "encoding/json"
    "net/http"

    "phishing-platform/database"
    "phishing-platform/internal/jwt"
)

func LoginHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req LoginRequest

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

	storedPassword,
		role,
		err :=
		database.GetUserByUsername(
			req.Username,
		)

	if err != nil {

		http.Error(
			w,
			"Invalid username",
			http.StatusUnauthorized,
		)

		return
	}

	if storedPassword != req.Password {

		http.Error(
			w,
			"Invalid password",
			http.StatusUnauthorized,
		)

		return
	}

	token, err :=
		jwt.GenerateToken(
			req.Username,
			role,
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