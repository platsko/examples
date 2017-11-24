package types

import (
	json "github.com/json-iterator/go"

	"lab/go-rest-api/app/repo"
)

type (
	// Client describes response element.
	Client struct {
		MobilePhone     string  `json:"mobilePhone"`
		Email           string  `json:"email,omitempty"`
		ValidationCode  string  `json:"validationCode"`
		AvailableAmount float64 `json:"availableAmount"`
	}
)

// @TODO fill Email, ValidationCode
// NewClient returns client element.
func NewClient(repo *repo.ClientRepo) *Client {
	balance, _ := repo.Balance()
	return &Client{
		MobilePhone:     repo.Phone,
		AvailableAmount: float64(balance) / 100,
	}
}

// MarshalJSON satisfies marshaler interface.
func (t *Client) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}
