package multidict

import (
	base "github.com/dploop/gostl/collections/associative_containers/tree/avl"
	"github.com/dploop/gostl/iterators"
	"github.com/dploop/gostl/traits"
	"github.com/dploop/gostl/types"
)

var _ iterators.MutableBidirectionalIterator = Iterator{}

type Iterator struct {
	base base.Iterator
}

func (i Iterator) Write(data types.Data) {
	i.base.Write(data)
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
	i.base = i.base.ImplNext()
	return i
}

func (i Iterator) Equal(other traits.EqualityComparable) bool {
	return i.ImplEqual(other.(Iterator))
}

func (i Iterator) ImplEqual(other Iterator) bool {
	return i == other
}

func (i Iterator) Read() types.Data {
	return i.base.Read()
}

func (i Iterator) Prev() traits.Decrementable {
	return i.ImplPrev()
}

func (i Iterator) ImplPrev() Iterator {
	i.base = i.base.ImplPrev()
	return i
}
