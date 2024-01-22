package pagination

import (
	"errors"
	"net/url"
	"strconv"
)

type Meta struct {
	// CurrentPage defaults to 1
	CurrentPage int `json:"current_page"`
	// Total defaults to 0
	Total int `json:"total"`
	// LastPage defaults to Total / Limit, otherwise 0
	LastPage int `json:"last_page"`
	// Limit defaults to 10
	Limit int `json:"limit"`
	// Offset defaults to Limit * CurrentPage, otherwise 0
	Offset int `json:"offset"`
}

// NewMeta creates a new instance for pagination meta
func NewMeta(q url.Values) (*Meta, error) {
	var err error
	pg := &Meta{
		CurrentPage: 1,
		Limit:       10,
	}

	if q.Get("page") != "" {
		pg.CurrentPage, err = strconv.Atoi(q.Get("page"))
		if err != nil || pg.CurrentPage <= 0 {
			return pg, errors.New("Incorrect query parameters: page")
		}
	}

	if q.Get("total") != "" {
		pg.Total, err = strconv.Atoi(q.Get("total"))
		if err != nil || pg.Total < 0 {
			return pg, errors.New("Incorrect query parameters: total")
		}
	}

	if q.Get("limit") != "" {
		pg.Limit, err = strconv.Atoi(q.Get("limit"))
		if err != nil || pg.Limit <= 0 {
			return pg, errors.New("Incorrect query parameters: limit")
		}
	}

	pg.Offset = (pg.CurrentPage - 1) * pg.Limit

	pg.SetNewTotal(pg.Total)

	return pg, nil
}

// SetNewTotal will set the total rows and the last page if the value is not 0
func (pgMeta *Meta) SetNewTotal(total int) {
	if total == 0 {
		return
	}

	pgMeta.Total = total
	pgMeta.LastPage = total / pgMeta.Limit
}
