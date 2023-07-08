package user

const (
	ProviderClient   = "client"
	ProviderAdmin    = "admin"
	ProviderInternal = "internal"
	ProviderGuest    = "guest"
)

type Identity interface {
	Provider() string
	ID() string
}

type identityState struct {
	provider string
	id       string
}

func (s identityState) Provider() string {
	return s.provider
}

func (s identityState) ID() string {
	return s.id
}
