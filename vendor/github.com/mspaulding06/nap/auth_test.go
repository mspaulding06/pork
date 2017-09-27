package nap

import "testing"

func TestAuthToken(t *testing.T) {
	token := NewAuthToken("somerandomtokenstring")
	header := token.AuthorizationHeader()
	if header != "token somerandomtokenstring" {
		t.Fail()
	}
}

func TestAuthBasic(t *testing.T) {
	basic := NewAuthBasic("user", "passw0rd")
	header := basic.AuthorizationHeader()
	if header != "basic dXNlcjpwYXNzdzByZA==" {
		t.Fail()
	}
}
