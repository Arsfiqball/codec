package something

// Queryable fields are fields that can be used in the where clause.
var queryableFields = []string{
	"id",
	"author_id",
	"name",
	"description",
}

// Groupable fields are fields that can be used in the group clause.
var groupableFields = []string{
	"name",
	"description",
}

// Sortable fields are fields that can be used in the sort clause.
var sortableFields = []string{
	"id",
	"name",
	"description",
}

// Withable fields are fields that can be used in the with clause.
var withableFields = []string{
	"author",
}

// Accumulable fields are fields that can be used in the accumulator.
var accumulableFields = []string{
	"count",
}

// isQueryableField returns true if the field is queryable.
func isGroupableField(field string) bool {
	for _, f := range groupableFields {
		if f == field {
			return true
		}
	}

	return false
}

// isSortableField returns true if the field is sortable.
func isSortableField(field string) bool {
	for _, f := range sortableFields {
		if f == field {
			return true
		}
	}

	return false
}

// isWithableField returns true if the field is withable.
func isWithableField(field string) bool {
	for _, f := range withableFields {
		if f == field {
			return true
		}
	}

	return false
}

// isAccumulableField returns true if the field is accumulable.
func isAccumulableField(field string) bool {
	for _, f := range accumulableFields {
		if f == field {
			return true
		}
	}

	return false
}
