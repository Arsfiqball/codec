package domain

import (
	"errors"
	"fmt"

	"github.com/Arsfiqball/talkback"
)

type Query struct {
	talkback.Query
	Search string
}

func (q Query) Validate() error {
	for _, c := range q.Conditions {
		if !c.Valid() || !isQueryableField(c.Field) {
			return errors.New(fmt.Sprintf("invalid condition for %s", c.Field))
		}
	}

	for _, s := range q.Sort {
		if !isSortableField(s.Field) {
			return errors.New(fmt.Sprintf("invalid sort for %s", s.Field))
		}
	}

	for _, a := range q.Accumulator {
		if !isAccumulableField(a) {
			return errors.New(fmt.Sprintf("invalid accumulator for %s", a))
		}
	}

	for _, g := range q.Group {
		if !isGroupableField(g) {
			return errors.New(fmt.Sprintf("invalid group for %s", g))
		}
	}

	for _, w := range q.With {
		if !isWithableField(w) {
			return errors.New(fmt.Sprintf("invalid with for %s", w))
		}
	}

	return nil
}
