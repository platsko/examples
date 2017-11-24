package middleware

import (
	"net/http"
)

// AuthRequest is a middleware for authorized request.
func AuthRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rs http.ResponseWriter, rq *http.Request) {
		if !authorized(rq) {
			http.Error(rs, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		next.ServeHTTP(rs, rq)
	})
}

// @FIXME there is a stub.
// @TODO implementation instead.
// authorized restricts access for a request.
func authorized(rq *http.Request) bool {
	token := rq.Header.Get("Authorization")
	if token != "" {
		b := make([]byte, 0)
		if _, err := rq.Body.Read(b); err != nil {
			return false
		}
	}

	return true
}
