package singly

import (
	"github.com/dploop/gostl/iterators"
	"github.com/dploop/gostl/traits"
	"github.com/dploop/gostl/types"
)

var _ iterators.MutableForwardIterator = Iterator{}

type Iterator struct {
	n *node
}

func (i Iterator) Write(data types.Data) {
	i.n.data = data
}

func (i Iterator) Clone() traits.Cloneable {
	return i.ImplClone()
}

func (i Iterator) ImplClone() Iterator {
	return i
}

func (i Iterator) Next() traits.Incrementable {
	return i.ImplNext()
}

func (i Iterator) ImplNext() Iterator {
	i.n = i.n.next
	return i
}

func (i Iterator) Equal(other traits.EqualityComparable) bool {
	return i.ImplEqual(other.(Iterator))
}

func (i Iterator) ImplEqual(other Iterator) bool {
	return i == other
}

func (i Iterator) Read() types.Data {
	return i.n.data
}
