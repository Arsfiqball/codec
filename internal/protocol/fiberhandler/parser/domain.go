package parser

import (
	"encoding/json"
	"feature/internal/value/domain"
	"net/url"
	"strconv"
)

func DomainPatchJSON(data []byte, patch *domain.Patch) error {
	var dto struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	patch.Name = domain.NewOmittable(dto.Name)
	patch.Email = domain.NewOmittable(dto.Email)
	patch.Password = domain.NewOmittable(dto.Password)

	return nil
}

func DomainQueryString(data []byte, query *domain.Query) error {
	values, err := url.ParseQuery(string(data))
	if err != nil {
		return err
	}

	limit, err := strconv.Atoi(values.Get("limit"))
	if err != nil {
		return err
	}

	skip, err := strconv.Atoi(values.Get("skip"))
	if err != nil {
		return err
	}

	var sort []domain.Sort

	for _, v := range values["sort"] {
		if v != "" && v[0] == '-' {
			sort = append(sort, domain.NewSort(v[1:], true))
		} else {
			sort = append(sort, domain.NewSort(v, false))
		}
	}

	*query = domain.NewQueryWithData(
		nil,
		values.Get("search"),
		limit,
		skip,
		sort,
		values["accumulator"],
		values["group"],
		values["with"],
	)

	return nil
}
