package something

// Stat is a stat of something.
type Stat struct {
	ID          string `json:"id"`          // This is a queryable field.
	Name        string `json:"name"`        // This is a queryable field.
	Description string `json:"description"` // This is a queryable field.
	Count       int    `json:"count"`       // This is a accumulative field.
}
