package request

import (
	"fmt"
	"net/http"

	"lab/go-rest-api/app/entity"
	"lab/go-rest-api/app/errors"
	t "lab/go-rest-api/app/repo/types"
	"lab/go-rest-api/database"
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
		Payment         Payment    `json:"payment"` // payment type
		Shift           uint64     `json:"shift" validate:"required"` // shift number
		Number          uint64     `json:"number" validate:"required"` // document number
		Uid             string     `json:"uid" validate:"required"` // document uid
		ReferenceUid    string     `json:"referenceUid"` // document reference uid
		DiscardedAmount string     `json:"discardedAmount"` // rounding of amounts
		Payments        []Payment  `json:"payments"` // payments list
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
	DocCalculate          DocAction = "calculate"
	DocPayByBonus         DocAction = "payByBonus"
	DocCancelBonusPayment DocAction = "cancelBonusPayment"
	DocConfirm            DocAction = "confirm"
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

	return r.extendValidate()
}

// @TODO make sure the implementation is correct.
// extendValidate validates request with extended rules.
func (r *documentRequest) extendValidate() error {
	switch f := r.form; f.Action {
	case DocCalculate:
		if err := validPositionsNonEmpty(f.Positions); err != nil {
			return err
		}

	case DocPayByBonus:
		if err := validPayModeNonFiscal(f.Payment.Mode); err != nil {
			return err
		}

	case DocCancelBonusPayment:
		if err := validPayModeNonFiscal(f.Payment.Mode); err != nil {
			return err
		}
		if err := validPaymentsNonEmpty(f.Payments); err != nil {
			return err
		}

	case DocConfirm:
		if err := validDocType(f.Type); err != nil {
			return err
		}
		if err := validPaymentsNonEmpty(f.Payments); err != nil {
			return err
		}
		if err := validPositionsNonEmpty(f.Positions); err != nil {
			return err
		}
		if err := validDocReferenceUid(f); err != nil {
			return err
		}

	default:
		return errors.ErrNotSupported()

	}

	return nil
}

// validDocType validates the document type.
func validDocType(docType t.DocType) error {
	switch docType {
	case t.DocTypeReceipt, t.DocTypeRefundReceipt:
		return nil
	}
	return fmt.Errorf("validate: unsupported document type '%s'", docType)
}

// validDocReferenceUid validates the document reference uid.
func validDocReferenceUid(f *DocumentForm) error {
	switch f.Type {
	case t.DocTypeReceipt: // this type has no reference
		return nil

	case t.DocTypeRefundReceipt: // try to read reference document
		e := entity.NewDocument(database.Orm())
		e.Uid = f.ReferenceUid
		if err := e.Read(); err == nil {
			return nil
		}
	}

	return fmt.Errorf("validate: reference document not found '%s'", f.ReferenceUid)
}

// validPayModeNonFiscal validates the payment mode is non-fiscal.
func validPayModeNonFiscal(payMode t.PayMode) error {
	if payMode == t.PayModeNonFiscal {
		return nil
	}
	return fmt.Errorf("validate: unsupported payment mode '%s'", payMode)
}

// validPaymentsNonEmpty validates payment list is not empty.
func validPaymentsNonEmpty(payments []Payment) error {
	if len(payments) > 0 {
		return nil
	}
	return fmt.Errorf("validate: payments cannot be empty")
}

// validPositionsNonEmpty validates position list is not empty.
func validPositionsNonEmpty(positions []Position) error {
	if len(positions) > 0 {
		return nil
	}
	return fmt.Errorf("validate: positions cannot be empty")
}
