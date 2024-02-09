package config

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	os.Setenv("API_PORT", "8080")

	config := Default()

	if reflect.TypeOf(config).String() != "*config.AppConfig" {
		t.Error("Default() did not get the correct type, wanted *config.AppConfig")
	}

	if config.Port != 8080 {
		t.Error("Default().Port expecting 8080, but got", config.Port)
	}

	if config.AccessTokenExp != time.Duration(15*time.Minute) {
		t.Error("Default().AccessTokenExp expecting time.Duration of 15 minutes, but got", config.AccessTokenExp.String())
	}
}

func TestAppConfig_WithProductionMode(t *testing.T) {
	config := Default().WithProductionMode(true)

	if reflect.TypeOf(config).String() != "*config.AppConfig" {
		t.Error("WithProductionMode(true) did not get the correct type, wanted *config.AppConfig")
	}

	if config.InProduction != true {
		t.Error("WithProductionMode(true) expecting InProduction to be true", config.InProduction)
	}
}
