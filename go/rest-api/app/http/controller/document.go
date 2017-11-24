package controller

import (
	"net/http"

	"github.com/go-chi/render"

	"lab/go-rest-api/app/errors"
	api "lab/go-rest-api/app/http/request"
	resp "lab/go-rest-api/app/http/respond"
	"lab/go-rest-api/app/repo"
	t "lab/go-rest-api/app/repo/types"
	service "lab/go-rest-api/app/service/document"
)

// Document controls requests by document route.
func Document(rs http.ResponseWriter, rq *http.Request) {
	request := api.DocumentRequest(rq)
	if err := request.Validate(); err != nil {
		resp.ErrResource(rs, rq, err)
		return
	}

	switch request.Action() {
	case api.DocCalculate:
		docCalculate(rs, request)
	case api.DocPayByBonus:
		docPayByBonus(rs, request)
	case api.DocCancelBonusPayment:
		docCancelBonusPayment(rs, request)
	case api.DocConfirm:
		docConfirm(rs, request)
	default:
		err := errors.ErrNotSupported()
		render.Respond(rs, rq, resp.ErrResponse(err))
	}
}

// docCalculate controls calculate action.
func docCalculate(rs http.ResponseWriter, rq api.DocumentRequester) {
	f := rq.Form()
	rp, err := repo.CreateDocument(&f)
	if err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to calculate document
	if err = service.Calculate(rp); err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to commit changes
	if err = rp.Transaction(func() error {
		if err = rp.Client.ClientBonus.Update(); err != nil {
			return err
		}
		return rp.Update()
	}); err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	resp.CalculateResource(rs, rq.HttpRequest(), rp)
}

// docPayByBonus controls pay-by-bonus action.
func docPayByBonus(rs http.ResponseWriter, rq api.DocumentRequester) {
	f := rq.Form()
	rp, err := repo.FetchDocument(&f)
	if err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to make payment
	if err = service.PayByBonus(rp, &f); err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to commit changes
	if err = rp.Transaction(func() error {
		if err = rp.Client.ClientBonus.Update(); err != nil {
			return err
		}
		return rp.Update()
	}); err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	resp.PayByBonusResource(rs, rq.HttpRequest(), rp)
}

// docCancelBonusPayment controls cancel-bonus-payment action.
func docCancelBonusPayment(rs http.ResponseWriter, rq api.DocumentRequester) {
	f := rq.Form()
	rp, err := repo.CreateDocument(&f)
	if rp == nil || err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to confirm payment
	if err = service.CancelBonusPayment(rp, &f); err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to commit changes
	if err = rp.Transaction(func() error {
		return rp.Update()
	}); err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	resp.CancelBonusPaymentResponse(rs, rq.HttpRequest(), rp)
}

// docConfirm controls confirm action.
func docConfirm(rs http.ResponseWriter, rq api.DocumentRequester) {
	var (
		f   = rq.Form()
		rp  *repo.DocumentRepo
		err error
	)

	switch f.Type {
	case t.DocTypeReceipt:
		rp, err = repo.FetchDocument(&f)
	case t.DocTypeRefundReceipt:
		rp, err = repo.CreateDocument(&f)
	default:
		err = errors.ErrNotSupported()
	}
	if err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to confirm document
	if err = service.Confirm(rp, &f); err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to commit changes
	if err = rp.Transaction(func() error {
		return rp.Update()
	}); err != nil {
		resp.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	resp.ConfirmResponse(rs, rq.HttpRequest(), rp)
}
