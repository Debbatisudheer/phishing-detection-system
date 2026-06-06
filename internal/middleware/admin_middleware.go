package middleware

import "net/http"

func AdminMiddleware(
	next http.HandlerFunc,
) http.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		next.ServeHTTP(
			w,
			r,
		)
	}
}