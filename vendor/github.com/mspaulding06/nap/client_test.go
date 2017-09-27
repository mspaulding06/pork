package nap

import (
	"net/http"
	"testing"
)

func TestProcessRequest(t *testing.T) {
	client := NewClient()
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		return nil
	})
	res := NewResource("/get", "GET", router)
	if err := client.ProcessRequest("https://httpbin.org", res, nil); err != nil {
		t.Fail()
	}
}
