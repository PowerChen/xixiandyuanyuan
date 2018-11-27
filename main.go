/*
this application is for XiXi and YuanYuan play
1.first XiXi or YuanYuan select a game, Then they played
2.find who win
3.repet 1 and 2
*/
package main

import (
	"fmt"
	"xixiandyuanyuan/models/concurrent"
	"xixiandyuanyuan/models/threeoptions"
	"xixiandyuanyuan/models/twonums"
)

func main() {
	fmt.Println("Hello XiXi and YuanYuan, Let's Player Gmaes!")
	twonums.Twonums()
	threeoptions.ThreeOptions()

	calacList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myresult := make(chan int, 2)
	go concurrent.Sum(calacList[:len(calacList)/2], myresult)
	go concurrent.Sum(calacList[len(calacList)/2:], myresult)

	sum1, sum2 := <-myresult, <-myresult

	fmt.Println("result:", sum1, sum2, sum1+sum2)
}
