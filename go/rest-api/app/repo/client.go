package repo

import (
	"lab/go-rest-api/app/entity"
	"lab/go-rest-api/app/errors"
	"lab/go-rest-api/app/http/request"
)

type (
	// ClientRepo wraps the client entity.
	ClientRepo struct {
		*entity.Client
	}
)

// Client returns a new client entity repository
// for specified client request form data.
func Client(form *request.ClientForm) *ClientRepo {
	return &ClientRepo{Client: entity.NewClient(form)}
}

// @FIXME there is a stub.
// @TODO implementation instead.
// Balance returns amount of bonuses on client balance.
func (r *ClientRepo) Balance() (int64, error) {
	return 0, errors.ErrNotImplemented()
}
