package repo

import (
	"lab/go-rest-api/app/entity"
	"lab/go-rest-api/app/http/request"
	"lab/go-rest-api/app/repo/types"
	"lab/go-rest-api/database"
)

type (
	// DocumentRepo wraps entity.
	DocumentRepo struct {
		*entity.Document
	}
)

// NewDocument returns entity repository.
func NewDocument() *DocumentRepo {
	return &DocumentRepo{
		Document: entity.NewDocument(database.Orm()),
	}
}

// DocumentCalculate returns entity repository
// for calculate action and specified request form data.
func DocumentCalculate(f *request.DocumentForm) (*DocumentRepo, error) {
	// create document
	repo, err := newDocument(f)
	if err != nil {
		return nil, err
	}

	repo.Type = f.Type
	repo.Status = types.DocStatusPending

	return repo, nil
}

// DocumentConfirm returns entity repository
// for confirm action and specified request form data.
func DocumentConfirm(f *request.DocumentForm) (*DocumentRepo, error) {
	// fetch requested document
	d, err := fetchDocument(f)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// DocumentPayByBonus returns entity repository
// for pay-by-bonus action and specified request form data.
func DocumentPayByBonus(f *request.DocumentForm) (*DocumentRepo, error) {
	// fetch requested document
	d, err := fetchDocument(f)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// fetchDocument returns repository of requested document.
func fetchDocument(f *request.DocumentForm) (*DocumentRepo, error) {
	// fetch requested document
	repo, err := newDocument(f)
	if err != nil {
		return nil, err
	}
	if err = repo.Read(); err != nil {
		return nil, err
	}

	return repo, nil
}

// newDocument returns repository filled with general fields.
func newDocument(f *request.DocumentForm) (*DocumentRepo, error) {
	// fetch requested client account
	client, err := clientRepo(f)
	if err != nil {
		return nil, err
	}

	// create and fill document
	repo := NewDocument()
	repo.ClientID = client.ID
	repo.Shift = f.Shift
	repo.Number = f.Number
	repo.Uid = f.Uid
	repo.Type = f.Type

	return repo, nil
}

// clientRepo returns repository of requested client.
func clientRepo(f *request.DocumentForm) (*ClientRepo, error) {
	// fetch requested client account
	client, err := Client(f.ClientForm)
	if err != nil {
		return nil, err
	}
	if err = client.Read(); err != nil {
		return nil, err
	}

	return client, nil
}
