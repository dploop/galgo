package double_list

import (
	"unsafe"
)

type List struct {
	prev *node
	next *node
	size Size
}

type node struct {
	prev *node
	next *node
	data Data
}

func (l *List) sent() *node {
	return (*node)(unsafe.Pointer(l))
}

func (l *List) assert(p *node) {
	if p == l.sent() {
		_ = *(*node)(nil)
	}
}

func NewList() *List {
	l := &List{}
	s := l.sent()
	s.prev, s.next = s, s
	return l
}

func (l *List) Empty() bool {
	return l.size == 0
}

func (l *List) Size() Size {
	return l.size
}

func (l *List) PushFront(data Data) {
	p := &node{data: data}
	l.link(l.sent().next, p, p)
	l.size++
}

func (l *List) Front() Data {
	return l.sent().next.data
}

func (l *List) PopFront() {
	p := l.sent().next
	l.assert(p)
	l.unlink(p, p)
	l.size--
}

func (l *List) PushBack(data Data) {
	p := &node{data: data}
	l.link(l.sent(), p, p)
	l.size++
}

func (l *List) Back() Data {
	return l.sent().prev.data
}

func (l *List) PopBack() {
	p := l.sent().prev
	l.assert(p)
	l.unlink(p, p)
	l.size--
}

func (l *List) Clear() {
	s := l.sent()
	s.prev, s.next = s, s
	l.size = 0
}

func (l *List) Begin() Iterator {
	return Iterator{l.sent().next}
}

func (l *List) End() Iterator {
	return Iterator{l.sent()}
}

func (l *List) ReverseBegin() Iterator {
	return Iterator{l.sent().prev}
}

func (l *List) ReverseEnd() Iterator {
	return Iterator{l.sent()}
}

func (l *List) Insert(i Iterator, data Data) Iterator {
	p := &node{data: data}
	l.link(i.node, p, p)
	l.size++
	return Iterator{p}
}

func (l *List) Erase(i Iterator) Iterator {
	p := i.node
	l.assert(p)
	r := p.next
	l.unlink(p, p)
	l.size--
	return Iterator{r}
}

func (l *List) link(p, head, tail *node) {
	p.prev.next = head
	head.prev = p.prev
	p.prev = tail
	tail.next = p
}

func (l *List) unlink(head, tail *node) {
	head.prev.next = tail.next
	tail.next.prev = head.prev
}
