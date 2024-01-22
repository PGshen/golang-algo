/*
 * @Descripttion:
 * @version:
 * @Date: 2024-01-16 13:33:51
 * @LastEditTime: 2024-01-19 13:22:34
 */
package datastruct

import (
	"container/list"
	"fmt"
)

type linkedListDeque struct {
	data *list.List
}

/* 初始化双端队列 */
func newLinkedListDeque() *linkedListDeque {
	return &linkedListDeque{
		data: list.New(),
	}
}

func (s *linkedListDeque) PushFirst(value any) {
	s.data.PushFront(value)
}

func (s *linkedListDeque) PushLast(value any) {
	s.data.PushBack(value)
}

func (s *linkedListDeque) PopFirst() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	s.data.Remove(e)
	return e.Value
}

func (s *linkedListDeque) PopLast() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

func (s *linkedListDeque) peekFirst() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Front().Value
}

func (s *linkedListDeque) peekLast() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Back().Value
}

func (s *linkedListDeque) size() int {
	return s.data.Len()
}

func (s *linkedListDeque) isEmpty() bool {
	return s.data.Len() == 0
}

func (s *linkedListDeque) toList() *list.List {
	return s.data
}

func (s *linkedListDeque) Print() {
	p := s.data.Front()
	for {
		fmt.Printf("%v ", p.Value)
		if p == s.data.Back() {
			break
		}
		p = p.Next()
	}
	fmt.Println()
}
