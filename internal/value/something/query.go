package something

// Query represents a query.
type Query struct {
	Conditions  []Condition
	Group       []string
	Sort        []string
	With        []string
	Accumulator []string
	Skip        int
	Limit       int
}

// Valid returns true if the query is valid.
func (q Query) Valid() bool {
	if len(q.Conditions) == 0 {
		return false
	}

	for _, condition := range q.Conditions {
		if !condition.Valid() {
			return false
		}
	}

	for _, field := range q.Group {
		if !isGroupableField(field) {
			return false
		}
	}

	for _, field := range q.Sort {
		if !isSortableField(field) {
			return false
		}
	}

	for _, field := range q.With {
		if !isWithableField(field) {
			return false
		}
	}

	for _, field := range q.Accumulator {
		if !isAccumulableField(field) {
			return false
		}
	}

	if q.Skip < 0 {
		return false
	}

	if q.Limit < 1 {
		return false
	}

	return true
}
