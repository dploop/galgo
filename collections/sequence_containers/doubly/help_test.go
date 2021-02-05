package doubly_test

import (
	"github.com/dploop/gostl/collections/sequence_containers/doubly"
)

func newList() *doubly.List {
	return doubly.NewList()
}

func newIterator() doubly.Iterator {
	l := newList()
	l.PushBack(0)
	return l.Begin()
}
