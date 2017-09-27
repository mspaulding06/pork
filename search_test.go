package pork

import "testing"

func TestSearchByKeyword(t *testing.T) {
	repositoryList := SearchByKeyword([]string{"one", "two", "three"})
	if repositoryList[0] != "myrepository" {
		t.Fail()
	}
}
