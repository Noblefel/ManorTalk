package token

import (
	"testing"
	"time"
)

var secretKey = "test"

var details = Details{
	SecretKey: secretKey,
	UserId:    1,
	UniqueId:  "test-uid",
	Duration:  5 * time.Minute,
}

func TestGenerate(t *testing.T) {
	_, err := Generate(details)
	if err != nil {
		t.Error("Generate() expecting no error")
	}
}

var sampleToken, _ = Generate(details)
var expiredToken, _ = Generate(Details{
	SecretKey: "expired",
})

func TestParse(t *testing.T) {
	var tests = []struct {
		name           string
		token          string
		secretKey      string
		expectedUserId int
		isError        bool
	}{
		{"success", sampleToken, secretKey, details.UserId, false},
		{"invalid token", "ascpdjapjdsacpjcdpasjcdpaj", "", 0, true},
		{"fail secret key", sampleToken, "", 0, true},
		{"expired token", expiredToken, "expired", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td, err := Parse(tt.secretKey, tt.token)
			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got: %v", err)
			}

			if err == nil && tt.isError {
				t.Errorf("expecting error")
			}

			if td != nil && td.UserId != tt.expectedUserId {
				t.Errorf("want user id %d, got %d", tt.expectedUserId, td.UserId)
			}
		})
	}
}
