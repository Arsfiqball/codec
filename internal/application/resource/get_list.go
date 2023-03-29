package resource

import (
	"context"
	"feature/internal/value/something"
	"feature/internal/value/user"
)

// GetListInput is the input for the GetList method.
type GetListInput struct {
	User  user.Entity
	Query something.Query
}

// GetListOutput is the output for the GetList method.
type GetListOutput struct {
	Docs []something.Entity `json:"docs"`
}

// GetList gets a list of resources.
func (s *Service) GetList(ctx context.Context, input GetListInput) (GetListOutput, error) {
	if input.User.IsEmpty() {
		return GetListOutput{}, ErrUserRequired
	}

	// Set a default limit.
	if input.Query.Limit == 0 {
		input.Query.Limit = 10
	}

	if !input.Query.Valid() {
		return GetListOutput{}, ErrInvalidQuery
	}

	docs, err := s.somethingRepo.Aggregate(ctx, input.Query)
	if err != nil {
		return GetListOutput{}, err
	}

	return GetListOutput{Docs: docs}, nil
}
