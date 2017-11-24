// Copyright © 2020 The EVEN Lab Team

package respond

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"evenlab/go-priority-api/app/repo"
)

type (
	// clientResponse describes of client response.
	clientResponse struct {
		repo    *repo.ClientRepo
		Code    int    `json:"code"`
		Display string `json:"cashierInformation,omitempty"`
		Printer string `json:"printingInformation,omitempty"`
	}
)

// Render satisfies Renderer interface.
func (r *clientResponse) Render(_ http.ResponseWriter, rq *http.Request) error {
	render.Status(rq, r.Code)
	return nil
}

// ClientResource wraps error resource for respond.
func ClientResource(rs http.ResponseWriter, rq *http.Request, r *repo.ClientRepo) {
	balance, err := r.Balance()
	if err != nil {
		render.Respond(rs, rq, ErrResponse(err))
		return
	}

	render.Respond(rs, rq, &clientResponse{
		repo:    r,
		Display: fmt.Sprintf("Account %s has %d bonuses.", r.Phone, balance),
		Printer: fmt.Sprintf("Account %s has %d bonuses.", r.Phone, balance),
	})
}
