package vector_test

import (
	"github.com/dploop/gostl/collections/sequence_containers/vector"
)

func newList() *vector.List {
	return vector.NewList()
}

func newIterator() vector.Iterator {
	l := newList()
	l.PushBack(0)
	return l.Begin()
}
