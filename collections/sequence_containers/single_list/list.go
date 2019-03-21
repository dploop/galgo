package single_list

type List struct {
	head *Node
	size Size
}

type Node struct {
	data Data
	next *Node
}

func NewList() *List {
	return &List{}
}

func (l *List) Empty() bool {
	return l.head == nil
}

func (l *List) Size() Size {
	return l.size
}

func (l *List) PushFront(data Data) {
	l.head = &Node{data: data, next: l.head}
	l.size++
}

func (l *List) Front() Data {
	return l.head.data
}

func (l *List) PopFront() {
	l.head = l.head.next
	l.size--
}

func (l *List) Clear() {
	l.head, l.size = nil, 0
}

func (l *List) Begin() Iterator {
	return Iterator{p: l.head}
}

func (l *List) End() Iterator {
	return Iterator{p: nil}
}

func (l *List) InsertAfter(i Iterator, data Data) Iterator {
	i.p.next = &Node{data: data, next: i.p.next}
	l.size++
	return Iterator{p: i.p.next}
}

func (l *List) EraseAfter(i Iterator) Iterator {
	i.p.next = i.p.next.next
	l.size--
	return Iterator{p: i.p.next}
}
