package single_list

import (
	"github.com/dploop/galgo/basics"
	"github.com/dploop/galgo/iterators"
)

var _ iterators.MutableForwardIterator = (*Iterator)(nil)

type Iterator struct {
	*node
}

func (i Iterator) Clone() basics.Cloneable {
	return i.ImplClone()
}

func (i Iterator) ImplClone() Iterator {
	return i
}

func (i Iterator) Read() iterators.Data {
	return i.ImplRead()
}

func (i Iterator) ImplRead() Data {
	return i.data
}

func (i *Iterator) Next() {
	i.node = i.next
}

func (i Iterator) Equal(other basics.EqualityComparable) bool {
	return i.ImplEqual(other.(Iterator))
}

func (i Iterator) ImplEqual(other Iterator) bool {
	return i == other
}

func (i Iterator) Write(data iterators.Data) {
	i.ImplWrite(data)
}

func (i Iterator) ImplWrite(data Data) {
	i.data = data
}
