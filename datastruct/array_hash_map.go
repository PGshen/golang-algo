package datastruct

import "fmt"

type pair struct {
	key   int
	value any
}

type arrayHashMap struct {
	buckets []*pair
}

func newArrayHashMap() *arrayHashMap {
	buckets := make([]*pair, 100)
	return &arrayHashMap{
		buckets: buckets,
	}
}

func (a *arrayHashMap) hashFunc(key int) int {
	return key % 100
}

func (a *arrayHashMap) get(key int) (bool, any) {
	index := a.hashFunc(key)
	pair := a.buckets[index]
	if pair == nil {
		return false, nil
	}
	return true, pair.value
}

func (a *arrayHashMap) put(key int, value any) {
	pair := &pair{key: key, value: value}
	index := a.hashFunc(key)
	a.buckets[index] = pair
}

func (a *arrayHashMap) remove(key int) {
	index := a.hashFunc(key)
	a.buckets[index] = nil
}

func (a *arrayHashMap) pairSet() []*pair {
	var pairs []*pair
	for _, pair := range a.buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func (a *arrayHashMap) keySet() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

func (a *arrayHashMap) valueSet() []any {
	var values []any
	for _, pair := range a.buckets {
		if pair != nil {
			values = append(values, pair.value)
		}
	}
	return values
}

func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Printf("%d -> %v\n", pair.key, pair.value)
		}
	}
}
