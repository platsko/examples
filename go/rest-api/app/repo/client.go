package repo

import (
	"lab/go-rest-api/app/entity"
	"lab/go-rest-api/app/errors"
	"lab/go-rest-api/app/http/request"
	"lab/go-rest-api/database"
)

type (
	// ClientRepo wraps entity.
	ClientRepo struct {
		*entity.Client
	}
)

// NewClient returns entity repository.
func NewClient(e *entity.Client) *ClientRepo {
	if e == nil {
		e = entity.NewClient(database.Orm())
	}
	return &ClientRepo{Client: e}
}

// Client returns entity repository
// for specified request form data.
func Client(form *request.ClientForm) (*ClientRepo, error) {
	repo := NewClient(nil)
	repo.Client.Phone = form.Login.Phone.Value()

	return repo, nil
}

// @FIXME there is a stub.
// @TODO implementation instead.
// Balance returns amount of bonuses on client balance.
func (r *ClientRepo) Balance() (int64, error) {
	return 0, errors.ErrNotImplemented()
}
