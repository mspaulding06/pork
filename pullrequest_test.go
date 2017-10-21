package pork

import (
	"testing"

	"github.com/mspaulding06/nap"
)

func TestPullRequest(t *testing.T) {
	token := "49117eb33240d82724587351e54434122667b3f9"
	GitHubAPI().SetAuth(nap.NewAuthToken(token))
	destRepo = "mspaulding06/testrepo:master"
	sourceRepo = "mspaulding06:mychanges"
	pullRequestTitle = "test pull request"
	pullRequestMessage = "here it is"
	if err := CreatePullRequest(); err != nil {
		t.Fail()
	}
}
