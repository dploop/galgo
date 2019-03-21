package single_list

import (
	"github.com/dploop/galgo/basics"
	"github.com/dploop/galgo/iterators"
)

var _ iterators.MutableForwardIterator = (*Iterator)(nil)

type Iterator struct {
	p *Node
}

func (i Iterator) Clone() basics.Cloneable {
	return i.ImplClone()
}

func (i Iterator) ImplClone() Iterator {
	return Iterator{p: i.p}
}

func (i Iterator) Read() iterators.Data {
	return i.ImplRead()
}

func (i Iterator) ImplRead() Data {
	return i.p.data
}

func (i *Iterator) Next() {
	i.p = i.p.next
}

func (i Iterator) Equal(other basics.EqualityComparable) bool {
	return i.ImplEqual(other.(Iterator))
}

func (i Iterator) ImplEqual(other Iterator) bool {
	return i.p == other.p
}

func (i Iterator) Write(data iterators.Data) {
	i.ImplWrite(data)
}

func (i Iterator) ImplWrite(data Data) {
	i.p.data = data
}
