package nap

import (
	"net/http"
	"testing"
)

func TestAPICall(t *testing.T) {
	api := NewAPI("https://httpbin.org")
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		return nil
	})
	res := NewResource("/get", "GET", router)
	api.AddResource("get", res)
	if err := api.Call("get", nil); err != nil {
		t.Fail()
	}
	resources := api.ResourceNames()
	if len(resources) != 1 || resources[0] != "get" {
		t.Fail()
	}
}

func TestAPIAuth(t *testing.T) {
	api := NewAPI("https://httpbin.org")
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		return nil
	})
	res := NewResource("/basic-auth/{{.user}}/{{.pass}}", "GET", router)
	api.AddResource("basicauth", res)
	api.SetAuth(&AuthBasic{
		Username: "user",
		Password: "passw0rd",
	})
	if err := api.Call("basicauth", map[string]string{
		"user": "user",
		"pass": "passw0rd",
	}); err != nil {
		t.Fail()
	}
}
