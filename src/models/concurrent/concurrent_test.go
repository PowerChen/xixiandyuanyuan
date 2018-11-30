package concurrent

import "testing"

func TestSum(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	result := make(chan int, 1)

	Sum(list, result)
	ret := <-result
	if ret != 15 {
		t.Errorf("sum failed got %v expected %v", ret, 15)
	}
}
