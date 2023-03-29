package user

// Entity represents a user.
type Entity struct {
	ID    string
	Roles []string
}

// HasRole returns true if the user has the role.
func (e Entity) HasRole(role string) bool {
	for _, r := range e.Roles {
		if r == role {
			return true
		}
	}

	return false
}

// IsEmpty returns true if the entity is empty.
func (e Entity) IsEmpty() bool {
	return e.ID == ""
}
