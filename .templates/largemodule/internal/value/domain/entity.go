package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Entity struct {
	id       string
	name     string
	email    string
	password string
}

func (e Entity) ID() string {
	return e.id
}

func (e Entity) Name() string {
	return e.name
}

func (e Entity) Email() string {
	return e.email
}

func (e Entity) Password() string {
	return e.password
}

func (e *Entity) Patch(p Patch) error {
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

func (e Entity) Validate() error {
	errs := validation.Errors{
		"id":       validation.Validate(e.id, validation.Required, is.UUIDv4),
		"name":     validation.Validate(e.name, validation.Required),
		"email":    validation.Validate(e.email, validation.Required, is.Email),
		"password": validation.Validate(e.password, validation.Required),
	}

	return errs.Filter()
}

func (e *Entity) ResetPassword() {
	e.password = ""
}
