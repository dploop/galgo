package initial_test

import (
	"github.com/dploop/gostl/collections/sequence_containers/initial"
	"github.com/dploop/gostl/types"
)

func newList(datas ...types.Data) *initial.List {
	return initial.NewList(datas...)
}

func newIterator() initial.Iterator {
	l := newList(0)
	return l.Begin()
}
