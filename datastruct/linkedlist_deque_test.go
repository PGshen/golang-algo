/*
 * @Descripttion:
 * @version:
 * @Date: 2024-01-16 13:42:36
 * @LastEditTime: 2024-01-19 13:22:09
 */
package datastruct

import (
	"fmt"
	"testing"
)

func TestDeque(t *testing.T) {
	deque := newLinkedListDeque()
	deque.PushFirst(1)
	deque.PushFirst(2)
	deque.PushLast("2222")
	fmt.Println(deque.PeekFirst())
	fmt.Println(deque.PeekLast())
	deque.PopFirst()
	deque.PushFirst("3333")
	deque.Print()
	deque.PopLast()
	deque.Print()

	fmt.Printf("%d", deque.Size())
}
