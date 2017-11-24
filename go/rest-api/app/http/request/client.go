package request

import (
	"fmt"
	"net/http"

	"lab/go-rest-api/app/repo/types"
)

type (
	// ClientRequester describes request interface.
	ClientRequester interface {
		ApiRequester      // use requester interface
		Validator         // use validator interface
		Form() ClientForm // Form returns form data

	}

	// ClientForm describes client request form.
	ClientForm struct {
		*ApiForm             // use api form fields
		Login    clientLogin `json:"client" validate:"required"`
	}

	// clientRequest describes client request.
	clientRequest struct {
		rq   *apiRequest
		form *ClientForm
	}

	// clientLogin wraps allowed identifiers for client login.
	clientLogin struct {
		Email string      `json:"email" validate:"email,required_without=mobilePhone"`
		Phone types.Phone `json:"mobilePhone" validate:"len=12,required_without=email"`
	}
)

var (
	// Make sure the type satisfies interface.
	_ ClientRequester = (*clientRequest)(nil)
)

// ClientRequest returns request interface.
func ClientRequest(request *http.Request) ClientRequester {
	return &clientRequest{
		rq:   &apiRequest{Request: request},
		form: &ClientForm{},
	}
}

// Form satisfies client request interface.
func (r *clientRequest) Form() ClientForm {
	form := *r.form
	return form
}

// HttpRequest satisfies request interface.
func (r *clientRequest) HttpRequest() *http.Request {
	return r.rq.Request
}

// Validate satisfies api request interface.
func (r *clientRequest) Validate() error {
	// try to parse request form data
	if err := r.rq.ParseForm(r.form); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	// validate request form data
	if err := validate.Struct(r.form); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}
