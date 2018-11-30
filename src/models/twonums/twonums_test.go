package twonums

import "testing"

//! Test
func TestTwonums(t *testing.T) {
	if !Twonums() {
		t.Log("no pass the test")
		t.Fail()
	}
}
