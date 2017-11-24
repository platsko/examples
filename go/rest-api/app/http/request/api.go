package request

import (
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// ApiRequester describes request interface.
	ApiRequester interface {
		HttpRequest() *http.Request
	}

	// Validator describes validation interface.
	Validator interface {
		Validate() error
	}

	// ApiForm describes the general fields of an api request form.
	ApiForm struct {
		Time    string `json:"dateTime" validate:"required"`
		OrgName string `json:"organization" validate:"required"`
		OrgUnit string `json:"businessUnit" validate:"required"`
		UnitPos string `json:"workPlace" validate:"required"`
	}

	// apiRequest wraps http request.
	apiRequest struct {
		*http.Request
	}
)

var (
	// validate keeps forms validator instance.
	validate = validator.New()
)

// ParseForm tries parse and decode specified form data.
func (r *apiRequest) ParseForm(form interface{}) error {
	// make sure the request is not empty
	if r.ContentLength == 0 {
		return fmt.Errorf("empty")
	}

	// try to parse request form data
	if err := r.Request.ParseForm(); err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	// try to decode request form data
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	return nil
}
