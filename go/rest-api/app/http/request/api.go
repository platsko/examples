package request

import (
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// Validator describes an api request interface.
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

	// apiRequest wraps the http request.
	apiRequest struct {
		*http.Request
	}
)

var (
	// validate keeps forms validator instance.
	validate = validator.New()
)

// ParseForm tries parse and decode specified form data.
func (api *apiRequest) ParseForm(form interface{}) error {
	// Make sure the request is not empty.
	if api.ContentLength == 0 {
		return fmt.Errorf("empty")
	}

	// Try to parse request form data.
	if err := api.Request.ParseForm(); err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	// Try to decode request form data.
	if err := json.NewDecoder(api.Body).Decode(form); err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	return nil
}
