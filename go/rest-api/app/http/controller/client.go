package controller

import (
	"net/http"

	api "lab/go-rest-api/app/http/request"
	"lab/go-rest-api/app/http/respond"
	"lab/go-rest-api/app/repo"
)

// FetchClient fetches or creates a client account.
func FetchClient(rs http.ResponseWriter, rq *http.Request) {
	request := api.ClientRequest(rq)
	if err := request.Validate(); err != nil {
		respond.ErrResource(rs, rq, err)
		return
	}

	form := request.Form()
	client, err := repo.Client(&form)
	if err != nil {
		respond.ErrResource(rs, rq, err)
		return
	}

	// fetch requested client account
	if err = client.Create(); err != nil {
		respond.ErrResource(rs, rq, err)
		return
	}

	respond.ClientResource(rs, rq, client)
}
