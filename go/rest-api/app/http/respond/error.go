package respond

import (
	"net/http"

	"github.com/go-chi/render"

	"lab/go-rest-api/app/errors"
)

type (
	// errResponse describes error response.
	errResponse struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	}
)

// Render satisfies renderer interface.
func (e *errResponse) Render(_ http.ResponseWriter, rq *http.Request) error {
	render.Status(rq, e.Code)
	return nil
}

// ErrResponse returns renderer interface for specified error.
func ErrResponse(err error) render.Renderer {
	return &errResponse{
		Code:  errors.ErrCodeUnknown,
		Error: err.Error(),
	}
}

// ErrResponse wraps error resource to respond.
func ErrResource(rs http.ResponseWriter, rq *http.Request, err error) {
	render.Respond(rs, rq, ErrResponse(err))
}
