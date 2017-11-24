package controller

import (
	"net/http"

	"github.com/go-chi/render"

	"lab/go-rest-api/app/errors"
	api "lab/go-rest-api/app/http/request"
	"lab/go-rest-api/app/http/respond"
	"lab/go-rest-api/app/repo"
	service "lab/go-rest-api/app/service/document"
)

// Document controls requests by document route.
func Document(rs http.ResponseWriter, rq *http.Request) {
	request := api.DocumentRequest(rq)
	if err := request.Validate(); err != nil {
		respond.ErrResource(rs, rq, err)
		return
	}

	switch request.Action() {
	case api.DocCalculate:
		docCalculate(rs, request)

	case api.DocConfirm:
		docConfirm(rs, request)

	case api.DocPayByBonus:
		docPayByBonus(rs, request)

	default:
		err := errors.ErrNotSupported()
		render.Respond(rs, rq, respond.ErrResponse(err))
	}
}

// docCalculate controls calculate bonus action.
func docCalculate(rs http.ResponseWriter, rq api.DocumentRequester) {
	f := rq.Form()
	rp, err := repo.DocumentCalculate(&f)
	if err != nil {
		respond.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	if err = rp.Create(); err != nil {
		respond.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	respond.DocumentResource(rs, rq.HttpRequest(), rp)
}

// docConfirm controls confirm bonus action.
func docConfirm(rs http.ResponseWriter, rq api.DocumentRequester) {
	f := rq.Form()
	rp, err := repo.DocumentConfirm(&f)
	if rp == nil || err != nil {
		respond.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to confirm payment
	if err = service.Confirm(rp, &f); err != nil {
		respond.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to commit changes
	if err = rp.Transaction(func() error {
		return rp.Update()
	}); err != nil {
		respond.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	respond.DocumentResource(rs, rq.HttpRequest(), rp)
}

// docPayByBonus controls payment by bonus action.
func docPayByBonus(rs http.ResponseWriter, rq api.DocumentRequester) {
	f := rq.Form()
	rp, err := repo.DocumentPayByBonus(&f)
	if err != nil {
		respond.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to make payment
	if err = service.PayByBonus(rp, &f); err != nil {
		respond.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	// try to commit changes
	if err = rp.Transaction(func() error {
		if err = rp.Client.ClientBonus.Update(); err != nil {
			return err
		}
		return rp.Update()
	}); err != nil {
		respond.ErrResource(rs, rq.HttpRequest(), err)
		return
	}

	respond.DocumentResource(rs, rq.HttpRequest(), rp)
}
