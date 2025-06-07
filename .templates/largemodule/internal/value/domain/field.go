package domain

const (
	FieldID          = "id"
	FieldName        = "name"
	FieldEmail       = "email"
	FieldPassword    = "password"
	FieldCount       = "count"
	FieldWithProfile = "profile"
)

var QueryableFields = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
}

var GroupableFields = []string{
	FieldName,
}

var SortableFields = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
}

var WithableFields = []string{
	FieldWithProfile,
}

var AccumulableFields = []string{
	FieldCount,
}

func isQueryableField(field string) bool {
	for _, f := range QueryableFields {
		if f == field {
			return true
		}
	}

	return false
}

func isGroupableField(field string) bool {
	for _, f := range GroupableFields {
		if f == field {
			return true
		}
	}

	return false
}

func isSortableField(field string) bool {
	for _, f := range SortableFields {
		if f == field {
			return true
		}
	}

	return false
}

func isWithableField(field string) bool {
	for _, f := range WithableFields {
		if f == field {
			return true
		}
	}

	return false
}

func isAccumulableField(field string) bool {
	for _, f := range AccumulableFields {
		if f == field {
			return true
		}
	}

	return false
}
