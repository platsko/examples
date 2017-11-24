package request

import (
	"fmt"
	"net/http"

	t "lab/go-rest-api/app/repo/types"
)

type (
	// DocumentRequester describes request interface.
	DocumentRequester interface {
		ApiRequester        // use requester interface
		Validator           // use validator interface
		Action() DocAction  // Action returns document action
		Form() DocumentForm // Form returns form data
	}

	// DocumentForm describes a document request form.
	DocumentForm struct {
		*ClientForm                // use client form fields
		Action          DocAction  `json:"action" validate:"required"` // action name
		Type            t.DocType  `json:"type" validate:"required"` // document type
		Payment         Payment    `json:"payment" validate:"required"` // payment type
		Shift           uint64     `json:"shift" validate:"required"` // shift number
		Number          uint64     `json:"number" validate:"required"` // document number
		Uid             string     `json:"uid" validate:"required"` // document uid
		ReferenceUid    string     `json:"referenceUid"` // document reference uid
		DiscardedAmount string     `json:"discardedAmount"` // rounding of amounts
		Payments        []Payment  `json:"payments" validate:"required"` // payments list
		Positions       []Position `json:"positions" validate:"required"` // positions list
	}

	// DocAction describes document's action.
	DocAction string

	// documentRequest describes a document request struct.
	documentRequest struct {
		rq   *apiRequest
		form *DocumentForm
	}
)

const (
	DocCalculate  DocAction = "calculate"
	DocConfirm    DocAction = "confirm"
	DocPayByBonus DocAction = "payByBonus"
)

var (
	// Make sure request satisfies requester interface.
	_ DocumentRequester = (*documentRequest)(nil)
)

// DocumentRequest returns request interface.
func DocumentRequest(request *http.Request) DocumentRequester {
	return &documentRequest{
		rq:   &apiRequest{Request: request},
		form: &DocumentForm{},
	}
}

// Action satisfies request interface.
func (r *documentRequest) Action() DocAction {
	return r.form.Action
}

// Form satisfies request interface.
func (r *documentRequest) Form() DocumentForm {
	return *r.form
}

// HttpRequest satisfies request interface.
func (r *documentRequest) HttpRequest() *http.Request {
	return r.rq.Request
}

// Validate satisfies request interface.
func (r *documentRequest) Validate() error {
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
