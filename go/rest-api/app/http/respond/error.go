package respond

import (
	"net/http"

	"github.com/go-chi/render"

	"lab/go-rest-api/app/errors"
)

type (
	// errResponse describes response.
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
// Can get specified error code as the second parameter of type int,
// otherwise, an error code of value unknown will be used: -1.
func ErrResponse(err error, code ...interface{}) render.Renderer {
	errCode := errors.ErrCodeUnknown
	if code != nil {
		if c, ok := code[0].(int); ok {
			errCode = c
		}
	}

	return &errResponse{
		Code:  errCode,
		Error: err.Error(),
	}
}

// ErrResponse wraps resource to respond.
func ErrResource(rs http.ResponseWriter, rq *http.Request, err error) {
	render.Respond(rs, rq, ErrResponse(err))
}
