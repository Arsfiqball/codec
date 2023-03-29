package resource

import (
	"context"
	"feature/internal/value/something"
	"feature/internal/value/user"

	"github.com/google/uuid"
)

// CreateInput is the input for the Create method.
type CreateInput struct {
	User  user.Entity
	Patch something.Patch
}

// CreateOutput is the output for the Create method.
type CreateOutput struct {
	Doc something.Entity `json:"doc"`
}

// Create creates a new resource.
func (s *Service) Create(ctx context.Context, input CreateInput) (CreateOutput, error) {
	if input.User.IsEmpty() {
		return CreateOutput{}, ErrUserRequired
	}

	entity := something.Entity{}
	entity.ID = uuid.NewString()
	entity.AuthorID = input.User.ID

	entity.Patch(input.Patch)

	err := entity.Validate()
	if err != nil {
		return CreateOutput{}, err
	}

	err = s.somethingRepo.Save(ctx, entity)
	if err != nil {
		return CreateOutput{}, err
	}

	return CreateOutput{Doc: entity}, nil
}
