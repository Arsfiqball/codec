package user

import "errors"

func NewEntity(
	provider string,
	id string,
) (Identity, error) {
	if provider == "" || id == "" {
		return nil, errors.New("user identity is empty")
	}

	if provider != ProviderClient &&
		provider != ProviderAdmin &&
		provider != ProviderInternal &&
		provider != ProviderGuest {
		return nil, errors.New("user identity provider is invalid")
	}

	return &identityState{
		provider: provider,
		id:       id,
	}, nil
}
