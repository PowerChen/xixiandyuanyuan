package threeoptions

import "testing"

//! Test
func TestThreeOptions(t *testing.T) {
	if !ThreeOptions() {
		t.Log("no pass the test")
		t.Fail()
	}
}
