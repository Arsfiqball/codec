package domain

type Stat interface {
	Fields() []string
	Count() int
}

type statState struct {
	fields []string
	count  int
}

func (s statState) Fields() []string {
	return s.fields
}

func (s statState) Count() int {
	return s.count
}
