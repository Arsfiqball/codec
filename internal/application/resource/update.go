package resource

import (
	"context"
	"feature/internal/value/something"
	"feature/internal/value/user"
)

// UpdateInput is the input for the Update method.
type UpdateInput struct {
	User  user.Entity
	Query something.Query
	Patch something.Patch
}

// UpdateOutput is the output for the Update method.
type UpdateOutput struct {
	Doc something.Entity `json:"doc"`
}

// Update updates a resource.
func (s *Service) Update(ctx context.Context, input UpdateInput) (UpdateOutput, error) {
	if input.User.IsEmpty() {
		return UpdateOutput{}, ErrUserRequired
	}

	if !input.User.HasRole("admin") {
		return UpdateOutput{}, ErrUserNotAdmin
	}

	input.Query.Limit = 1

	entities, err := s.somethingRepo.Aggregate(ctx, input.Query)
	if err != nil {
		return UpdateOutput{}, err
	}

	if len(entities) == 0 {
		return UpdateOutput{}, ErrResourceNotFound
	}

	entity := entities[0]

	if entity.AuthorID != input.User.ID {
		return UpdateOutput{}, ErrUserNotAuthor
	}

	entity.Patch(input.Patch)

	err = entity.Validate()
	if err != nil {
		return UpdateOutput{}, err
	}

	err = s.somethingRepo.Save(ctx, entity)
	if err != nil {
		return UpdateOutput{}, err
	}

	return UpdateOutput{Doc: entity}, nil
}
