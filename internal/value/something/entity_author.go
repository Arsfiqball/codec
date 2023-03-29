package something

import "reflect"

// EntityAuthor is the author of something.
type EntityAuthor struct {
	ID   string `json:"id"`   // This is a queryable field.
	Name string `json:"name"` // This is a queryable field.
}

// Validate validates the entity.
func (e EntityAuthor) Validate() error {
	if e.ID == "" {
		return ErrIDRequired
	}

	if e.Name == "" {
		return ErrNameRequired
	}

	return nil
}

// IsEmpty returns true if the entity is empty.
func (e EntityAuthor) IsEmpty() bool {
	return reflect.ValueOf(e).IsZero()
}
