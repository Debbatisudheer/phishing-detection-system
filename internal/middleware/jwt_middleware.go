package middleware

import (
	"net/http"
	"strings"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(
	"phishing-platform-secret",
)

func JWTMiddleware(
	next http.HandlerFunc,
) http.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		authHeader :=
			r.Header.Get(
				"Authorization",
			)

		if authHeader == "" {

			http.Error(
				w,
				"Missing token",
				http.StatusUnauthorized,
			)

			return
		}

		tokenString :=
			strings.TrimPrefix(
				authHeader,
				"Bearer ",
			)

		token, err :=
			jwtlib.Parse(
				tokenString,
				func(
					token *jwtlib.Token,
				) (interface{}, error) {

					return SecretKey,
						nil
				},
			)

		if err != nil {

			println(
				"JWT ERROR:",
				err.Error(),
			)

			http.Error(
				w,
				"Invalid token",
				http.StatusUnauthorized,
			)

			return
		}

		if !token.Valid {

			println(
				"TOKEN NOT VALID",
			)

			http.Error(
				w,
				"Invalid token",
				http.StatusUnauthorized,
			)

			return
		}

		println(
			"TOKEN VALID",
		)

		next.ServeHTTP(
			w,
			r,
		)
	}
}