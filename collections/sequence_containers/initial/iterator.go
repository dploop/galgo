package initial

import (
	"github.com/dploop/gostl/iterators"
	"github.com/dploop/gostl/traits"
	"github.com/dploop/gostl/types"
)

var _ iterators.InputIterator = Iterator{}

type Iterator struct {
	l *List
	n types.Size
}

func (i Iterator) Write(data types.Data) {
	i.l.slice[i.n] = data
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
	i.n++
	return i
}

func (i Iterator) Equal(other traits.EqualityComparable) bool {
	return i.ImplEqual(other.(Iterator))
}

func (i Iterator) ImplEqual(other Iterator) bool {
	return i == other
}

func (i Iterator) Read() types.Data {
	return i.l.slice[i.n]
}
