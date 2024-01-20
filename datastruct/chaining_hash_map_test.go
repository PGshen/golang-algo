package datastruct

import (
	"fmt"
	"testing"
)

func TestChainingHashMap(t *testing.T) {
	hashMap := newChainingHashMap()
	hashMap.put(1, "11")
	hashMap.put(1, "22")
	hashMap.put(5, "333")
	hashMap.print()
	hashMap.put(22, "111")
	hashMap.put(23, "222")
	hashMap.print()

	fmt.Println(hashMap.get(22))

	hashMap.remove(1)
	hashMap.print()
}
