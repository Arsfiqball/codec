package domain

type Stat struct {
	name  string
	count int
}

func (s Stat) Name() string {
	return s.name
}

func (s Stat) Count() int {
	return s.count
}
