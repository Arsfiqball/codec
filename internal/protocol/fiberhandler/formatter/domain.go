package formatter

import "feature/internal/value/domain"

type DomainEntityDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DomainStatDTO struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func FormatDomainEntityDTO(ent domain.Entity) DomainEntityDTO {
	return DomainEntityDTO{
		ID:       ent.ID(),
		Name:     ent.Name(),
		Email:    ent.Email(),
		Password: ent.Password(),
	}
}

func FormatDomainStatDTO(stat domain.Stat) DomainStatDTO {
	return DomainStatDTO{
		Name:  stat.Name(),
		Count: stat.Count(),
	}
}
