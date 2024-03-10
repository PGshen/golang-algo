package doublepointer

import (
	"fmt"
	"testing"
)

func TestSlidingWindowMaximum(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
}
