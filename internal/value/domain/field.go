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
