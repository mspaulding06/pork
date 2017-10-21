package pork

import (
	"testing"

	"github.com/mspaulding06/nap"
)

func TestForkRepository(t *testing.T) {
	token := "49117eb33240d82724587351e54434122667b3f9"
	GitHubAPI().SetAuth(nap.NewAuthToken(token))
	if err := ForkRepository("mspaulding06/testrepo"); err != nil {
		t.Fail()
	}
}
