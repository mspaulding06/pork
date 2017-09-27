package nap

import (
	"net/http"
	"net/url"
	"testing"
)

func TestUnknownStatusCode(t *testing.T) {
	router := NewRouter()
	fakeURL, err := url.Parse("https://httpbin.org/doesnotexist")
	if err != nil {
		t.Fail()
	}
	resp := &http.Response{
		Request: &http.Request{
			URL: fakeURL,
		},
		StatusCode: 404,
	}
	if err := router.CallFunc(resp, nil); err == nil {
		t.Fail()
	}
}
