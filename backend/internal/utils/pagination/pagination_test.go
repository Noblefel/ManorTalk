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
		{"success", url.Values{"page": {"1"}, "total": {"20000"}}, false},
		{"success without query params", nil, false},
		{"error negative page", url.Values{"page": {"-1"}}, true},
		{"error invalid page", url.Values{"page": {"x"}}, true},
		{"error negative total", url.Values{"total": {"-1"}}, true},
		{"error invalid total", url.Values{"total": {"x"}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMeta(tt.q, 1)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error,  got: %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}

func TestMeta_SetNewTotal(t *testing.T) {
	t.Run("SetNewTotal 0", func(t *testing.T) {
		pgMeta, _ := NewMeta(url.Values{}, 1)
		pgMeta.SetNewTotal(0, 1)

		if pgMeta.Total != 0 || pgMeta.LastPage != 0 {
			t.Error("should not affect pgMeta.Total and pgMeta.LastPage")
		}
	})

	t.Run("SetNewTotal 100, 10", func(t *testing.T) {
		pgMeta, _ := NewMeta(url.Values{}, 1)
		pgMeta.SetNewTotal(100, 10)

		if pgMeta.Total != 100 || pgMeta.LastPage != 10 {
			t.Error("should set pgMeta.Total to 100 and pgMeta.LastPage to 10")
		}
	})
}
