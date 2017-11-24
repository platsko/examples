package entity

type (
	// OrmManager describes entity interface.
	OrmManager interface {
		// Set entity's database table name.
		TableName() string

		// Create creates a new entity.
		Create() error

		// Read reads an existing entity.
		Read() error

		// Create updates an existing entity.
		Update() error

		// Delete deletes an existing entity.
		Delete() error
	}
)
