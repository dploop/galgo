package double_list

type List struct {
	sent Node
	size Size
}

type Node struct {
	data Data
	prev *Node
	next *Node
}

func NewList() *List {
	l := &List{}
	sent := &l.sent
	sent.prev = sent
	sent.next = sent
	return l
}

func (l *List) Empty() bool {
	sent := &l.sent
	return sent.next == sent
}

func (l *List) Size() Size {
	return l.size
}

func (l *List) PushFront(data Data) {
	sent := &l.sent
	node := &Node{data: data, prev: sent, next: sent.next}
	node.prev.next, node.next.prev = node, node
	l.size++
}

func (l *List) Front() Data {
	sent := &l.sent
	return sent.next.data
}

func (l *List) PopFront() {
	sent := &l.sent
	node := sent.next
	if node == sent {
		panic("no such element")
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

func (l *List) PushBack(data Data) {
	sent := &l.sent
	node := &Node{data: data}
	node.prev = sent.prev
	sent.prev.next = node
	node.next = sent
	sent.prev = node
	l.size++
}

func (l *List) Back() Data {
	sent := &l.sent
	return sent.prev.data
}

func (l *List) PopBack() {
	sent := &l.sent
	node := sent.prev
	if node == sent {
		panic("no such element")
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

func (l *List) Clear() {
	sent := &l.sent
	sent.prev = sent
	sent.next = sent
	l.size = 0
}

func (l *List) Begin() Iterator {
	sent := &l.sent
	return Iterator{p: sent.next}
}

func (l *List) End() Iterator {
	sent := &l.sent
	return Iterator{p: sent}
}

func (l *List) ReverseBegin() Iterator {
	sent := &l.sent
	return Iterator{p: sent.prev}
}

func (l *List) ReverseEnd() Iterator {
	sent := &l.sent
	return Iterator{p: sent}
}

func (l *List) InsertBefore(i Iterator, data Data) Iterator {
	node := &Node{data: data, prev: i.p.prev, next: i.p}
	node.prev.next, node.next.prev = node, node
	l.size++
	return Iterator{p: node}
}

func (l *List) InsertAfter(i Iterator, data Data) Iterator {
	node := &Node{data: data, prev: i.p, next: i.p.next}
	node.prev.next, node.next.prev = node, node
	l.size++
	return Iterator{p: node}
}

func (l *List) EraseBefore(i Iterator) Iterator {
	node := i.p.prev
	if node == i.p {
		panic("invalid iterator")
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
	return Iterator{p: i.p.prev}
}

func (l *List) EraseAfter(i Iterator) Iterator {
	node := i.p.next
	if node == i.p {
		panic("invalid iterator")
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
	return Iterator{p: i.p.next}
}
