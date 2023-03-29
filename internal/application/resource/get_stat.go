package resource

import (
	"context"
	"feature/internal/value/something"
	"feature/internal/value/user"
)

// GetStatInput is the input for the GetStat method.
type GetStatInput struct {
	User  user.Entity
	Query something.Query
}

// GetStatOutput is the output for the GetStat method.
type GetStatOutput struct {
	Docs []something.Stat `json:"docs"`
}

// GetStat gets a list of stats.
func (s *Service) GetStat(ctx context.Context, input GetStatInput) (GetStatOutput, error) {
	if input.User.IsEmpty() {
		return GetStatOutput{}, ErrUserRequired
	}

	// Set a default limit.
	if input.Query.Limit == 0 {
		input.Query.Limit = 10
	}

	if !input.Query.Valid() {
		return GetStatOutput{}, ErrInvalidQuery
	}

	docs, err := s.somethingRepo.AggregateStats(ctx, input.Query)
	if err != nil {
		return GetStatOutput{}, err
	}

	return GetStatOutput{Docs: docs}, nil
}
