// Copyright © 2020 The EVEN Lab Team

package controller

import (
	"net/http"

	api "evenlab/go-priority-api/app/http/request"
	"evenlab/go-priority-api/app/http/respond"
	"evenlab/go-priority-api/app/repo"
)

// FetchClient fetches or creates a client account.
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
