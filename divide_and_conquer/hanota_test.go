package divideandconquer

import (
	"container/list"
	"fmt"
	"testing"
)

func TestHanota(t *testing.T) {
	A := list.New()
	B := list.New()
	C := list.New()
	A.PushBack("1")
	A.PushBack("2")
	A.PushBack("3")
	solveHanota(A, B, C)
	for e := C.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
