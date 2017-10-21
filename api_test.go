package pork

import "testing"

func TestGitHubAPI(t *testing.T) {
	if GitHubAPI() == nil {
		t.Fail()
	}
}
