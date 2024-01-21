package datastruct

import "fmt"

type pair3 struct {
	key   int
	value any
}

type openAddressingHashMap struct {
	size        int
	capacity    int
	loadThres   float64
	extendRatio int
	buckets     []pair3
	removed     pair3
}

func newOpenAddressingHashMap() *openAddressingHashMap {
	buckets := make([]pair3, 4)
	return &openAddressingHashMap{
		size:        0,
		capacity:    4,
		loadThres:   2.0 / 3.0,
		extendRatio: 2,
		buckets:     buckets,
		removed: pair3{
			key:   -1,
			value: nil,
		},
	}
}

func (m *openAddressingHashMap) hashFunc(key int) int {
	return key % m.capacity
}

func (m *openAddressingHashMap) loadFactor() float64 {
	return float64(m.size) / float64(m.capacity)
}

func (m *openAddressingHashMap) get(key int) any {
	idx := m.hashFunc(key)
	for i := 0; i < m.capacity; i++ {
		// 线性探测
		j := (idx + i) % m.capacity
		if m.buckets[j] == (pair3{}) {
			return nil
		}
		if m.buckets[j].key == key && m.buckets[j] != m.removed {
			return m.buckets[j].value
		}
	}
	return nil
}

func (m *openAddressingHashMap) put(key int, value any) {
	if m.loadFactor() > m.loadThres {
		m.extend()
	}
	idx := m.hashFunc(key)
	for i := 0; i < m.capacity; i++ {
		// 线性探测
		j := (idx + i) % m.capacity
		if m.buckets[j] == (pair3{}) || m.buckets[j] == m.removed {
			m.buckets[j] = pair3{
				key:   key,
				value: value,
			}
			m.size++
			return
		}
		// 存在则替换
		if m.buckets[j].key == key {
			m.buckets[j].value = value
			return
		}
	}
}

func (m *openAddressingHashMap) remove(key int) {
	idx := m.hashFunc(key)
	for i := 0; i < m.capacity; i++ {
		j := (idx + i) % m.capacity
		// 要删除的key不存在
		if m.buckets[j] == (pair3{}) {
			return
		}
		// 找到要删除的key，设置删除标记
		if m.buckets[j].key == key {
			m.buckets[j] = m.removed
			m.size--
		}
	}
}

func (m *openAddressingHashMap) extend() {
	tmpBuckets := make([]pair3, len(m.buckets))
	copy(tmpBuckets, m.buckets)

	m.capacity *= m.extendRatio
	m.buckets = make([]pair3, m.capacity)
	m.size = 0

	for _, p := range tmpBuckets {
		if p != (pair3{}) && p != m.removed {
			m.put(p.key, p.value)
		}
	}
}

func (m *openAddressingHashMap) print() {
	for _, p := range m.buckets {
		if p != (pair3{}) {
			fmt.Printf("%d -> %v\n", p.key, p.value)
		} else {
			fmt.Println("nil")
		}
	}
}
