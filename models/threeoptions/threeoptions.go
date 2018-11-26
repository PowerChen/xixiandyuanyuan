/*
game has three options
rock, paper,scissors
1.XiXi and YuanYuan choise one show to other
2.then rock win scissors, scissors win paper, paper win rock
*/
package threeoptions

import (
	"fmt"
	"testing"
)

func ThreeOptions() bool {
	fmt.Println("Ok, XiXi and YuanYuan begin play rock,paper,scissors!")
	return true
}

//! Test
func TestThreeOptions(t *testing.T) {
	if !ThreeOptions() {
		t.Log("no pass the test")
		t.Fail()
	}
}
