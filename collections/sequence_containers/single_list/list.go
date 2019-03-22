package single_list

type List struct {
	sent Node
	size Size
}

type Node struct {
	data Data
	next *Node
}

func NewList() *List {
	l := &List{size: 0}
	sent := &l.sent
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
	node := &Node{data: data, next: sent.next}
	sent.next = node
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
	sent.next = node.next
	l.size--
}

func (l *List) Clear() {
	sent := &l.sent
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

func (l *List) InsertAfter(i Iterator, data Data) Iterator {
	node := &Node{data: data, next: i.p.next}
	i.p.next = node
	l.size++
	return Iterator{p: i.p.next}
}

func (l *List) EraseAfter(i Iterator) Iterator {
	node := i.p.next
	if node == i.p {
		panic("invalid iterator")
	}
	i.p.next = node.next
	l.size--
	return Iterator{p: i.p.next}
}
