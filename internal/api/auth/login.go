package auth

import (
	"encoding/json"
	"net/http"

	userrepo "phishing-platform/database/users"
	"phishing-platform/internal/jwt"
)

func LoginHandler(
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

	// Validate input
	if req.Username == "" ||
		req.Password == "" {

		http.Error(
			w,
			"username and password are required",
			http.StatusBadRequest,
		)

		return
	}

	storedPassword,
		role,
		err :=
		userrepo.GetUserByUsername(
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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(w).Encode(
		map[string]string{
			"token": token,
		},
	)
}