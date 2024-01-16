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
		t.Error("Generate() should not return error")
	}
}

var sampleToken, _ = Generate(details)
var expiredToken, _ = Generate(Details{
	SecretKey: "expired",
})

var parseTests = []struct {
	name           string
	token          string
	secretKey      string
	expectedUserId int
	isError        bool
}{
	{
		name:           "parse-ok",
		token:          sampleToken,
		secretKey:      secretKey,
		expectedUserId: details.UserId,
		isError:        false,
	},
	{
		name:      "parse-fail-invalid-token",
		token:     "ascpdjapjdsacpjcdpasjcdpaj",
		secretKey: "",
		isError:   true,
	},
	{
		name:      "parse-fail-secret-key",
		token:     sampleToken,
		secretKey: "",
		isError:   true,
	},
	{
		name:      "parse-fail-expired-token",
		token:     expiredToken,
		secretKey: "expired",
		isError:   true,
	},
}

func TestParse(t *testing.T) {
	for _, tt := range parseTests {
		td, err := Parse(tt.secretKey, tt.token)
		if err != nil && !tt.isError {
			t.Errorf("%s should not return error: %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return some error", tt.name)
		}

		if td != nil && td.UserId != tt.expectedUserId {
			t.Errorf("%s returned incorrect user id, wanted %d got %d", tt.name, tt.expectedUserId, td.UserId)
		}
	}
}
