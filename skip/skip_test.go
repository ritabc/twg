package skip

import "testing"

func TestThing(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Log("this test ran!")
}
