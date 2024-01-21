package datastruct

import (
	"fmt"
	"testing"
)

func TestOpenAddressingHashMap(t *testing.T) {
	hashMap := newOpenAddressingHashMap()
	hashMap.put(1, "11")
	hashMap.put(5, "55")
	hashMap.put(2, "222")

	hashMap.print()

	fmt.Println(hashMap.get(2))

	hashMap.remove(5)
	hashMap.print()

	hashMap.put(3, "23232")
	hashMap.put(3, "33333")
	hashMap.put(4, "4")
	hashMap.print()
}
