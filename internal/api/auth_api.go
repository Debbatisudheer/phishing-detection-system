package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
	"phishing-platform/internal/jwt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

	err = database.CreateUser(
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