// Copyright © 2020 The EVEN Lab Team

package request

import (
	"fmt"
	"net/http"

	"evenlab/go-priority-api/app/repo/types"
)

type (
	// ClientIface describes a client request interface.
	ClientIface interface {
		// Use api request interface.
		ApiIface

		// Form returns the client form data.
		Form() *ClientForm
	}

	// ClientForm describes a client request form.
	ClientForm struct {
		// Use api form fields.
		ApiForm

		// Client's form fields.
		Client mobilePhone `json:"client" validate:"required"`
	}

	// clientRequest describes a client request struct.
	clientRequest struct {
		rq   *apiRequest
		form *ClientForm
	}

	// mobilePhone wraps the PhoneNumber type for the client form field.
	mobilePhone struct {
		PhoneNumber types.PhoneNumber `json:"mobilePhone" validate:"required,len=12"`
	}
)

var (
	// Make sure the clientRequest type satisfies api interface.
	_ ClientIface = (*clientRequest)(nil)
)

// ClientRequest returns a new client request interface
// for the specified http request.
func ClientRequest(request *http.Request) ClientIface {
	return &clientRequest{
		rq:   &apiRequest{Request: request},
		form: &ClientForm{},
	}
}

// Form satisfies client request interface.
func (r *clientRequest) Form() *ClientForm {
	return r.form
}

// Validate satisfies api request interface.
func (r *clientRequest) Validate() error {
	// Try to parse request form data.
	if err := r.rq.ParseForm(r.form); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	// Validate request form data.
	if err := validate.Struct(r.form); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}
