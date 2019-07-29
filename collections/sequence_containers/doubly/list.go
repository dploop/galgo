package doubly

import (
	"github.com/dploop/gostl/types"
)

type List struct {
	sent *node
	size types.Size
}

type node struct {
	prev *node
	next *node
	data types.Data
}

func (l *List) assert(p *node) {
	if p == l.sent {
		_ = *(*node)(nil)
	}
}

func NewList() *List {
	s := &node{}
	s.prev, s.next = s, s
	return &List{sent: s}
}

func (l *List) Size() types.Size {
	return l.size
}

func (l *List) Empty() bool {
	return l.Size() == 0
}

func (l *List) PushFront(data types.Data) {
	p := &node{data: data}
	l.link(l.sent.next, p, p)
	l.size++
}

func (l *List) Front() types.Data {
	return l.sent.next.data
}

func (l *List) PopFront() {
	p := l.sent.next
	l.assert(p)
	l.unlink(p, p)
	l.size--
}

func (l *List) PushBack(data types.Data) {
	p := &node{data: data}
	l.link(l.sent, p, p)
	l.size++
}

func (l *List) Back() types.Data {
	return l.sent.prev.data
}

func (l *List) PopBack() {
	p := l.sent.prev
	l.assert(p)
	l.unlink(p, p)
	l.size--
}

func (l *List) Clear() {
	s := l.sent
	s.prev, s.next = s, s
	l.size = 0
}

func (l *List) Begin() Iterator {
	return Iterator{n: l.sent.next}
}

func (l *List) End() Iterator {
	return Iterator{n: l.sent}
}

func (l *List) ReverseBegin() Iterator {
	return Iterator{n: l.sent.prev}
}

func (l *List) ReverseEnd() Iterator {
	return Iterator{n: l.sent}
}

func (l *List) Insert(i Iterator, data types.Data) Iterator {
	p := &node{data: data}
	l.link(i.n, p, p)
	l.size++
	return Iterator{n: p}
}

func (l *List) Erase(i Iterator) Iterator {
	p := i.n
	l.assert(p)
	r := p.next
	l.unlink(p, p)
	l.size--
	return Iterator{n: r}
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
