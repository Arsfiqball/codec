package parser

import (
	"encoding/json"
	"feature/internal/application/resource"
	"feature/internal/value/domain"
	"net/url"

	"github.com/Arsfiqball/talkback-lancer"
)

type DomainPatchDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto DomainPatchDTO) ToPatch() domain.Patch {
	return domain.Patch{
		Name:     domain.NewOmittable(dto.Name),
		Email:    domain.NewOmittable(dto.Email),
		Password: domain.NewOmittable(dto.Password),
	}
}

type DomainBulkOpsDTO struct {
	Ops []struct {
		Type  string         `json:"type"`
		Query string         `json:"query"`
		Patch DomainPatchDTO `json:"patch"`
	} `json:"ops"`
}

func (dto DomainBulkOpsDTO) ToBulkOps() ([]resource.Ops, error) {
	var ops []resource.Ops

	for _, v := range dto.Ops {
		var query domain.Query

		if err := DomainQueryString([]byte(v.Query), &query); err != nil {
			return nil, err
		}

		ops = append(ops, resource.Ops{
			Type:  v.Type,
			Query: query,
			Patch: v.Patch.ToPatch(),
		})
	}

	return ops, nil
}

func DomainPatchJSON(data []byte, patch *domain.Patch) error {
	var dto DomainPatchDTO

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	*patch = dto.ToPatch()

	return nil
}

func DomainQueryString(data []byte, query *domain.Query) error {
	values, err := url.ParseQuery(string(data))
	if err != nil {
		return err
	}

	rql, err := talkback.FromURLValues(values)
	if err != nil {
		return err
	}

	*query = domain.Query{
		Query:  rql,
		Search: values.Get("search"),
	}

	return nil
}

func DomainBulkOpsJSON(data []byte, ops *[]resource.Ops) error {
	var dto DomainBulkOpsDTO

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	nops, err := dto.ToBulkOps()
	if err != nil {
		return err
	}

	*ops = nops

	return nil
}
