package router

import (
	"reflect"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/service/auth"
	"github.com/Noblefel/ManorTalk/backend/internal/service/post"
	"github.com/Noblefel/ManorTalk/backend/internal/service/user"
)

func TestNewRouter(t *testing.T) {
	var c *config.AppConfig
	var as auth.AuthService
	var us user.UserService
	var ps post.PostService
	router := NewRouter(c, as, us, ps)

	typeString := reflect.TypeOf(router).String()
	if typeString != "*router.router" {
		t.Error("NewRouter() did not get the correct type, wanted *router.router")
	}
}

func TestRouter_Routes(t *testing.T) {
	var c *config.AppConfig
	var as auth.AuthService
	var us user.UserService
	var ps post.PostService
	router := NewRouter(c, as, us, ps)

	mux := router.Routes()

	typeString := reflect.TypeOf(mux).String()
	if typeString != "*chi.Mux" {
		t.Error("router Routes() did not get the correct type, wanted *chi.Mux")
	}
}
