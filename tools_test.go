package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools

	s := testTools.RandomString(10)

	if len(s) != 10 {
		t.Errorf("Random string length expected to be %d but is %d.", 10, len(s))
	}
}
