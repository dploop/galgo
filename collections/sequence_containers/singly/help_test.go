package singly_test

import (
	"github.com/dploop/gostl/collections/sequence_containers/singly"
)

func newList() *singly.List {
	return singly.NewList()
}

func newIterator() singly.Iterator {
	l := newList()
	l.PushFront(0)
	return l.Begin()
}
