package resource

import (
	"context"
	"feature/internal/value/something"
	"feature/internal/value/user"
)

// DeleteInput is the input for the Delete method.
type DeleteInput struct {
	User  user.Entity
	Query something.Query
}

// DeleteOutput is the output for the Delete method.
type DeleteOutput struct {
	Doc something.Entity `json:"doc"`
}

// Delete deletes a resource.
func (s *Service) Delete(ctx context.Context, input DeleteInput) (DeleteOutput, error) {
	if input.User.IsEmpty() {
		return DeleteOutput{}, ErrUserRequired
	}

	if !input.User.HasRole("admin") {
		return DeleteOutput{}, ErrUserNotAdmin
	}

	input.Query.Limit = 1

	entities, err := s.somethingRepo.Aggregate(ctx, input.Query)
	if err != nil {
		return DeleteOutput{}, err
	}

	if len(entities) == 0 {
		return DeleteOutput{}, ErrResourceNotFound
	}

	entity := entities[0]

	if entity.AuthorID != input.User.ID {
		return DeleteOutput{}, ErrUserNotAuthor
	}

	err = s.somethingRepo.DeleteByID(ctx, entity.ID)
	if err != nil {
		return DeleteOutput{}, err
	}

	return DeleteOutput{Doc: entity}, nil
}
