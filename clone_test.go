package pork

import "testing"

func TestCloneRepository(t *testing.T) {
	if err := CloneRepository("myrepository", "", false); err != nil {
		t.Fail()
	}
}
