package datastruct

import (
	"fmt"
	"strings"
)

type pair2 struct {
	key   int
	value any
}

type chainingHashMap struct {
	size        int
	capacity    int
	loadThres   float64
	extendRatio int
	buckets     [][]pair2
}

func newChainingHashMap() *chainingHashMap {
	buckets := make([][]pair2, 4)
	for i := 0; i < 4; i++ {
		buckets[i] = make([]pair2, 0)
	}
	return &chainingHashMap{
		size:        0,
		capacity:    4,
		loadThres:   2.0 / 3.0,
		extendRatio: 2,
		buckets:     buckets,
	}
}

func (m *chainingHashMap) hashFunc(key int) int {
	return key % m.capacity
}

func (m *chainingHashMap) loadFactor() float64 {
	return float64(m.size) / float64(m.capacity)
}

func (m *chainingHashMap) get(key int) any {
	idx := m.hashFunc(key)
	bucket := m.buckets[idx]
	for _, p := range bucket {
		if p.key == key {
			return p.value
		}
	}
	return nil
}

func (m *chainingHashMap) put(key int, value any) {
	if m.loadFactor() > m.loadThres {
		m.extend()
	}
	idx := m.hashFunc(key)
	for i := range m.buckets[idx] {
		if m.buckets[idx][i].key == key {
			m.buckets[idx][i].value = value
			return
		}
	}
	p := pair2{
		key:   key,
		value: value,
	}
	m.buckets[idx] = append(m.buckets[idx], p)
	m.size++
}

func (m *chainingHashMap) remove(key int) {
	idx := m.hashFunc(key)
	for i, p := range m.buckets[idx] {
		if p.key == key {
			m.buckets[idx] = append(m.buckets[idx][:i], m.buckets[idx][i+1:]...)
			m.size--
			break
		}
	}
}

func (m *chainingHashMap) extend() {
	tmpBuckets := make([][]pair2, len(m.buckets))
	for i := 0; i < len(m.buckets); i++ {
		tmpBuckets[i] = make([]pair2, len(m.buckets[i]))
		copy(tmpBuckets[i], m.buckets[i])
	}

	m.capacity *= m.extendRatio
	m.buckets = make([][]pair2, m.capacity)
	for i := 0; i < m.capacity; i++ {
		m.buckets[i] = make([]pair2, 0)
	}
	m.size = 0

	for _, bucket := range tmpBuckets {
		for _, p := range bucket {
			m.put(p.key, p.value)
		}
	}
}

func (m *chainingHashMap) print() {
	var builder strings.Builder

	for _, bucket := range m.buckets {
		builder.WriteString("[")
		for _, p := range bucket {
			builder.WriteString(fmt.Sprintf("%d -> %v; ", p.key, p.value))
		}
		builder.WriteString("]")
		fmt.Println(builder.String())
		builder.Reset()
	}
	fmt.Println()
}
