package pagination

import (
	"net/url"
	"testing"
)

func TestNewMeta(t *testing.T) {
	tests := []struct {
		name    string
		q       url.Values
		isError bool
	}{
		{
			name: "newMeta-ok",
			q: url.Values{
				"page":  {"1"},
				"total": {"20000"},
			},
			isError: false,
		},
		{
			name:    "newMeta-ok-2",
			q:       nil,
			isError: false,
		},
		{
			name: "newMeta-error-page",
			q: url.Values{
				"page": {"-1"},
			},
			isError: true,
		},
		{
			name: "newMeta-error-page-2",
			q: url.Values{
				"page": {"x"},
			},
			isError: true,
		},
		{
			name: "newMeta-error-total",
			q: url.Values{
				"total": {"-1"},
			},
			isError: true,
		},
		{
			name: "newMeta-error-total-2",
			q: url.Values{
				"total": {"x"},
			},
			isError: true,
		},
	}

	for _, tt := range tests {
		_, err := NewMeta(tt.q, 1)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error: %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return some error", tt.name)
		}
	}
}

func TestMeta_SetNewTotal(t *testing.T) {
	pgMeta, _ := NewMeta(url.Values{}, 1)
	pgMeta.SetNewTotal(0, 1)

	if pgMeta.Total != 0 || pgMeta.LastPage != 0 {
		t.Error("SetNewTotal(0) should not affect pgMeta.Total and pgMeta.LastPage")
	}

	pgMeta, _ = NewMeta(url.Values{}, 1)
	pgMeta.SetNewTotal(100, 10)

	if pgMeta.Total != 100 || pgMeta.LastPage != 10 {
		t.Error("SetNewTotal(100) should set pgMeta.Total to 100 and pgMeta.LastPage to 10")
	}
}
