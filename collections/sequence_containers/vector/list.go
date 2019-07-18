package vector

import (
	"github.com/dploop/gostl/types"
)

type List struct {
	array []types.Data
}

func NewList() *List {
	return &List{}
}

func (l *List) Size() types.Size {
	return len(l.array)
}

func (l *List) Empty() bool {
	return l.Size() == 0
}

func (l *List) Get(i types.Size) types.Data {
	return l.array[i]
}

func (l *List) Set(n types.Size, data types.Data) {
	l.array[n] = data
}

func (l *List) PushBack(data types.Data) {
	l.array = append(l.array, data)
}

func (l *List) Back() types.Data {
	return l.array[len(l.array)-1]
}

func (l *List) PopBack() {
	l.array = l.array[:len(l.array)-1]
}


func (l *List) Clear() {
	l.array = nil
}

func (l *List) Begin() Iterator {
	return Iterator{l: l, n: 0}
}

func (l *List) End() Iterator {
	return Iterator{l: l, n: len(l.array)}
}

func (l *List) ReverseBegin() Iterator {
	return Iterator{l: l, n: len(l.array) - 1}
}

func (l *List) ReverseEnd() Iterator {
	return Iterator{l: l, n: -1}
}
