package pork

import "testing"

func TestForkRepository(t *testing.T) {
	if err := ForkRepository("myrepository"); err != nil {
		t.Fail()
	}
}
