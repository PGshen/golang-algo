package datastruct

import (
	"fmt"
	"testing"
)

func TestArrayHashMap(t *testing.T) {
	hashMap := newArrayHashMap()
	hashMap.put(1, "hah")
	hashMap.put(2, "kkk")
	hashMap.put(3, "ddd")
	hashMap.print()
	fmt.Println(hashMap.get(2))
	hashMap.remove(2)
	hashMap.print()
	fmt.Printf("%v\n", hashMap.keySet())
	fmt.Printf("%v\n", hashMap.valueSet())
	fmt.Printf("%v\n", hashMap.pairSet())
}
