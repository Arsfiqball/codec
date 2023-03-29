package something

// Condition represents a condition.
type Condition struct {
	Field string
	Value string
	Op    string
}

// Valid returns true if the condition is valid.
func (c Condition) Valid() bool {
	if c.Field == "" || c.Value == "" || c.Op == "" {
		return false
	}

	for _, field := range queryableFields {
		if field == c.Field {
			return true
		}
	}

	return false
}
