/*
concurrent example
*/

package concurrent

func Sum(list []int, result chan int) {
	sum := 0
	for _, v := range list {
		sum += v
	}
	result <- sum
}
