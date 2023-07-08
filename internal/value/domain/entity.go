package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Entity interface {
	ID() string
	Name() string
	Email() string
	Password() string
	Patch(Patch) error
	Validate() error
}

type entityState struct {
	id       string
	name     string
	email    string
	password string
}

func (e entityState) ID() string {
	return e.id
}

func (e entityState) Name() string {
	return e.name
}

func (e entityState) Email() string {
	return e.email
}

func (e entityState) Password() string {
	return e.password
}

func (e *entityState) Patch(p Patch) error {
	if p.Name.Valid() {
		e.name = p.Name.Value()
	}

	if p.Email.Valid() {
		e.email = p.Email.Value()
	}

	if p.Password.Valid() {
		e.password = p.Password.Value()
	}

	return nil
}

func (e entityState) Validate() error {
	errs := validation.Errors{
		"id":       validation.Validate(e.id, validation.Required, is.UUIDv4),
		"name":     validation.Validate(e.name, validation.Required),
		"email":    validation.Validate(e.email, validation.Required, is.Email),
		"password": validation.Validate(e.password, validation.Required),
	}

	return errs.Filter()
}
