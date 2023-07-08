package domain

type Stat interface {
	Name() string
	Count() int
}

type statState struct {
	name  string
	count int
}

func (s statState) Name() string {
	return s.name
}

func (s statState) Count() int {
	return s.count
}
