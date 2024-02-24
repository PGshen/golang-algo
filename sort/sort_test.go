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

func TestQuickSort(t *testing.T) {
	nums := []int{2, 1, 19, 23, 12, 4, 9, 22}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestMergeSort(t *testing.T) {
	nums := []int{2, 1, 19, 23, 12, 4, 9, 22}
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestHeapSort(t *testing.T) {
	nums := []int{2, 1, 19, 23, 12, 4, 9, 22}
	heapSort(nums)
	fmt.Println(nums)
}

func TestBucketSort(t *testing.T) {
	nums := []int{2, 1, 19, 23, 12, 4, 9, 22, 23, 18, 2, 3, 33, 50, 199, 22, 7}
	sortedNums := bucketSort(nums)
	fmt.Println(sortedNums)
}

func Test1(t *testing.T) {
	fmt.Printf("%v", 3/2)
}
