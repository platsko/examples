package request

import (
	"fmt"
	"net/http"

	"lab/go-rest-api/app/repo/types"
)

type (
	// ClientValidator describes a client validator interface.
	ClientValidator interface {
		// Use api request interface.
		Validator

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
	// Make sure the client request type satisfies validator interface.
	_ ClientValidator = (*clientRequest)(nil)
)

// ClientRequest returns a new client validator interface
// for specified http request.
func ClientRequest(request *http.Request) ClientValidator {
	return &clientRequest{
		rq:   &apiRequest{Request: request},
		form: &ClientForm{},
	}
}

// Form satisfies client validator interface.
func (r *clientRequest) Form() *ClientForm {
	return r.form
}

// Validate satisfies validator interface.
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
