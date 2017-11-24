package repo

import (
	"lab/go-rest-api/app/entity"
	"lab/go-rest-api/app/http/request"
	t "lab/go-rest-api/app/repo/types"
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

// Balance returns amount of bonuses on client balance.
func (r *ClientRepo) Balance() *t.Amount {
	return t.NewAmount(r.Amount)
}
