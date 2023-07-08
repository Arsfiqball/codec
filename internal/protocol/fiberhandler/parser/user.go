package parser

import (
	"errors"
	"feature/internal/value/user"
	"strings"
)

func UserIdentityToken(token string, ent *user.Identity) error {
	s := strings.Split(token, ":")

	if len(s) != 2 {
		return errors.New("invalid user identity token")
	}

	nent, err := user.NewEntity(s[0], s[1])
	if err != nil {
		return err
	}

	*ent = nent

	return nil
}
