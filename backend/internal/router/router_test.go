package router

import (
	"reflect"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
)

func TestNewRouter(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	router := NewRouter(c, db)

	typeString := reflect.TypeOf(router).String()
	if typeString != "*router.router" {
		t.Error("NewRouter() did not get the correct type, wanted *router.router")
	}
}

func TestRouter_Routes(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	router := NewRouter(c, db)

	mux := router.Routes()

	typeString := reflect.TypeOf(mux).String()
	if typeString != "*chi.Mux" {
		t.Error("router Routes() did not get the correct type, wanted *chi.Mux")
	}
}
