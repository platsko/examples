package controller

import (
	"net/http"

	api "lab/go-rest-api/app/http/request"
	"lab/go-rest-api/app/http/respond"
	"lab/go-rest-api/app/repo"
)

// FetchClient fetches or creates a client entity.
func FetchClient(rs http.ResponseWriter, rq *http.Request) {
	request := api.ClientRequest(rq)
	if err := request.Validate(); err != nil {
		respond.ErrResource(rs, rq, err)
		return
	}

	// Fetch requested client account.
	client := repo.Client(request.Form())
	if err := client.Create(); err != nil {
		respond.ErrResource(rs, rq, err)
		return
	}

	respond.ClientResource(rs, rq, client)
}
