package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{2, 1, 19, 23, 12, 4, 9, 22}
	selectionSort(nums)
	fmt.Println(nums)
}

func TestBubbleSort(t *testing.T) {
	nums := []int{2, 1, 19, 23, 12, 4, 9, 22}
	bubbleSort(nums)
	fmt.Println(nums)
	nums2 := []int{2, 1, 19, 23, 12, 4, 9, 22}
	bubbleSortWithFlag(nums2)
	fmt.Println(nums2)
}

func TestInsertionSort(t *testing.T) {
	nums := []int{2, 1, 19, 23, 12, 4, 9, 22}
	insertionSort(nums)
	fmt.Println(nums)
}
