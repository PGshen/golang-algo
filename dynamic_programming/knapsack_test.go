package dynamicprogramming

import (
	"fmt"
	"testing"
)

func TestKnapsackDFS(t *testing.T) {
	wgt := []int{10, 20, 20}
	val := []int{60, 110, 120}
	cap := 30
	fmt.Println(knapsackDFS(wgt, val, 3, cap))
}

func TestKnapsackDFSMem(t *testing.T) {
	wgt := []int{10, 20, 20}
	val := []int{60, 110, 120}
	cap := 30
	l := len(wgt)
	mem := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		mem[i] = make([]int, cap+1)
		for j := 0; j <= cap; j++ {
			mem[i][j] = -1
		}
	}
	fmt.Println(knapsackDFSMem(wgt, val, mem, 3, cap))
}

func TestKnapsackDP(t *testing.T) {
	wgt := []int{10, 20, 20}
	val := []int{60, 110, 120}
	cap := 30
	fmt.Println(knapsackDP(wgt, val, cap))
}

func TestKnapsackDPComp(t *testing.T) {
	wgt := []int{10, 20, 20}
	val := []int{60, 110, 120}
	cap := 30
	fmt.Println(knapsackDPComp(wgt, val, cap))
}
