package pork

import "testing"

func TestCloneRepository(t *testing.T) {
	if err := CloneRepository("mspaulding06/testrepo", "master", false); err != nil {
		t.Fail()
	}
}
