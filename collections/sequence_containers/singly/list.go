package singly

import (
	"unsafe"

	"github.com/dploop/gostl/types"
)

type List struct {
	next *node
	size types.Size
}

type node struct {
	next *node
	data types.Data
}

func (l *List) sent() *node {
	return (*node)(unsafe.Pointer(l))
}

func NewList() *List {
	return &List{}
}

func (l *List) Size() types.Size {
	return l.size
}

func (l *List) Empty() bool {
	return l.Size() == 0
}

func (l *List) PushFront(data types.Data) {
	s := l.sent()
	s.next = &node{next: s.next, data: data}
	l.size++
}

func (l *List) Front() types.Data {
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
	return Iterator{n: l.sent().next}
}

func (l *List) End() Iterator {
	return Iterator{}
}

func (l *List) InsertAfter(i Iterator, data types.Data) Iterator {
	i.n.next = &node{next: i.n.next, data: data}
	l.size++
	return Iterator{n: i.n.next}
}

func (l *List) EraseAfter(i Iterator) Iterator {
	i.n.next = i.n.next.next
	l.size--
	return Iterator{n: i.n.next}
}
