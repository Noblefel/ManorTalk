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
	// Offset defaults to Limit * CurrentPage, otherwise 0
	Offset int `json:"offset"`
}

// NewMeta creates a new instance for pagination meta
func NewMeta(q url.Values, limit int) (*Meta, error) {
	var err error
	pg := &Meta{CurrentPage: 1}

	if q.Get("page") != "" {
		pg.CurrentPage, err = strconv.Atoi(q.Get("page"))
		if err != nil || pg.CurrentPage <= 0 {
			return pg, errors.New("incorrect query parameters: page")
		}
	}

	if q.Get("total") != "" {
		pg.Total, err = strconv.Atoi(q.Get("total"))
		if err != nil || pg.Total < 0 {
			return pg, errors.New("incorrect query parameters: total")
		}
	}

	pg.Offset = (pg.CurrentPage - 1) * limit

	pg.SetNewTotal(pg.Total, limit)

	return pg, nil
}

// SetNewTotal will set the total rows and the last page if the value is not 0
func (pgMeta *Meta) SetNewTotal(total, limit int) {
	if total == 0 {
		return
	}

	pgMeta.Total = total
	pgMeta.LastPage = total / limit
}
