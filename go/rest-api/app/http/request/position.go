package request

type (
	// Position describes position form field.
	Position struct {
		Index        uint64 `json:"index" validate:"required"`        // position index in document
		Id           string `json:"id" validate:"required"`           // position identity
		Text         string `json:"text" validate:"required"`         // position description
		Price        string `json:"price" validate:"required"`        // position price amount
		MinimumPrice string `json:"minimumPrice" validate:"required"` // position minimum price amount
		Quantity     string `json:"quantity" validate:"required"`     // quantity of position
		TotalAmount  string `json:"totalAmount" validate:"required"`  // position total amount
	}
)
