package greedy

import (
	"fmt"
	"testing"
)

func TestMaxCapacity(t *testing.T) {
	ht := []int{3, 8, 5, 2, 7, 7, 3, 4}
	fmt.Println(maxCapacity(ht))
}
