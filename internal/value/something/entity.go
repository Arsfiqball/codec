package something

import "reflect"

// Entity is the entity of something.
type Entity struct {
	ID          string `json:"id"`          // This is a queryable field.
	AuthorID    string `json:"author_id"`   // This is a queryable field.
	Name        string `json:"name"`        // This is a queryable field.
	Description string `json:"description"` // This is a queryable field.

	Author EntityAuthor `json:"author"` // This is a withable field.
}

// Validate validates the entity.
func (e Entity) Validate() error {
	if e.ID == "" {
		return ErrIDRequired
	}

	if e.Name == "" {
		return ErrNameRequired
	}

	if !e.Author.IsEmpty() {
		if err := e.Author.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// IsEmpty returns true if the entity is empty.
func (e Entity) IsEmpty() bool {
	return reflect.ValueOf(e).IsZero()
}

// Patch patches the entity.
func (e Entity) Patch(p Patch) Entity {
	if !p.Name.Omit {
		e.Name = p.Name.Value
	}

	if !p.Description.Omit {
		e.Description = p.Description.Value
	}

	return e
}
