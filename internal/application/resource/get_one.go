package resource

import (
	"context"
	"feature/internal/value/something"
	"feature/internal/value/user"
)

// GetOneInput is the input for the GetOne method.
type GetOneInput struct {
	User  user.Entity
	Query something.Query
}

// GetOneOutput is the output for the GetOne method.
type GetOneOutput struct {
	Doc something.Entity `json:"doc"`
}

func (s *Service) GetOne(ctx context.Context, input GetOneInput) (GetOneOutput, error) {
	if input.User.IsEmpty() {
		return GetOneOutput{}, ErrUserRequired
	}

	input.Query.Limit = 1

	if !input.Query.Valid() {
		return GetOneOutput{}, ErrInvalidQuery
	}

	docs, err := s.somethingRepo.Aggregate(ctx, input.Query)
	if err != nil {
		return GetOneOutput{}, err
	}

	if len(docs) == 0 {
		return GetOneOutput{}, ErrResourceNotFound
	}

	doc := docs[0]

	return GetOneOutput{Doc: doc}, nil
}
