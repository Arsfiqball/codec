package action

import (
	"context"
	"feature/internal/value/something"
	"feature/internal/value/user"
)

// Service is the application service for resource.
type ClaimSomethingInput struct {
	User user.Entity
	ID   string // The ID of the something to claim.
}

// Service is the application service for resource.
type ClaimSomethingOutput struct {
	Doc something.Entity `json:"doc"`
}

// ClaimSomething claims a something.
func (s *Service) ClaimSomething(ctx context.Context, input ClaimSomethingInput) (ClaimSomethingOutput, error) {
	if input.User.IsEmpty() {
		return ClaimSomethingOutput{}, ErrUserRequired
	}

	entity, err := s.somethingRepo.GetByID(ctx, input.ID)
	if err != nil {
		return ClaimSomethingOutput{}, err
	}

	if entity.IsEmpty() {
		return ClaimSomethingOutput{}, ErrResourceNotFound
	}

	entity.AuthorID = input.User.ID

	err = s.somethingRepo.Save(ctx, entity)
	if err != nil {
		return ClaimSomethingOutput{}, err
	}

	return ClaimSomethingOutput{Doc: entity}, nil
}
