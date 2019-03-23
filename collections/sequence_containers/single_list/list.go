package single_list

import (
	"unsafe"
)

type List struct {
	next *node
	size Size
}

type node struct {
	next *node
	data Data
}

func (l *List) sent() *node {
	return (*node)(unsafe.Pointer(l))
}

func NewList() *List {
	return &List{}
}

func (l *List) Empty() bool {
	return l.size == 0
}

func (l *List) Size() Size {
	return l.size
}

func (l *List) PushFront(data Data) {
	s := l.sent()
	s.next = &node{next: s.next, data: data}
	l.size++
}

func (l *List) Front() Data {
	return l.sent().next.data
}

func (l *List) PopFront() {
	s := l.sent()
	s.next = s.next.next
	l.size--
}

func (l *List) Clear() {
	l.sent().next = nil
	l.size = 0
}

func (l *List) Begin() Iterator {
	return Iterator{l.sent().next}
}

func (l *List) End() Iterator {
	return Iterator{}
}

func (l *List) InsertAfter(i Iterator, data Data) Iterator {
	i.next = &node{next: i.next, data: data}
	l.size++
	return Iterator{i.next}
}

func (l *List) EraseAfter(i Iterator) Iterator {
	i.next = i.next.next
	l.size--
	return Iterator{i.next}
}
