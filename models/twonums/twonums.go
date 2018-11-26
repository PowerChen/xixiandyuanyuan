/*
this game is for XiXi and YuanYuan
game ruler:
1.XiXi apply two num for YuanYuan to calculate
2.if YuanYuan's return is right then YuanYuan apply two num for XiXi calculate
3.repet 1 and 2
*/
package twonums

import (
	"fmt"
	"testing"
)

func Twonums() bool {
	fmt.Println("Ok XiXi and YuanYuan Now Play twonums Game")
	return true
}

//! Test
func TestTwonums(t testing.T) {
	if !Twonums() {
		t.Log("no pass the test")
		t.Fail()
	}
}
