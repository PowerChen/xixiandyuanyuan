/*
this application is for XiXi and YuanYuan play
1.first XiXi or YuanYuan select a game, Then they played
2.find who win
3.repet 1 and 2
*/
package main

import (
	"fmt"
	"xixiandyuanyuan/models/threeoptions"
	"xixiandyuanyuan/models/twonums"
)

func main() {
	fmt.Println("Hello XiXi and YuanYuan, Let's Player Gmaes!")
	twonums.Twonums()
	threeoptions.ThreeOptions()

}

func Even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
