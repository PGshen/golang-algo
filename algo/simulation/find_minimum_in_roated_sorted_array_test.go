package simulation

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 4}))
	fmt.Println(findMin([]int{4, 5, 6, -1, 0, 1, 4}))
}
