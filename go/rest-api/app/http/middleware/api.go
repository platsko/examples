// Copyright © 2020 The EVEN Lab Team

package middleware

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// ApiRequest is a common middleware for request.
func ApiRequest(handler http.Handler) http.Handler {
	// Use chi middleware section
	middleware.RequestID(handler)
	middleware.Logger(handler)
	middleware.Recoverer(handler)
	middleware.URLFormat(handler)
	render.SetContentType(render.ContentTypeJSON)

	return http.HandlerFunc(func(rs http.ResponseWriter, rq *http.Request) {
		handler.ServeHTTP(rs, rq)
	})
}
