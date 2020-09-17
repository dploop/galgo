package iterators

import (
	"github.com/dploop/gostl/constraints"
	"github.com/dploop/gostl/types"
)

type RandomAccessIterator interface {
	BidirectionalIterator
	constraints.LessThanComparable
	At(diff types.Size) types.Data
	Advance(diff types.Size) RandomAccessIterator
	Distance(other RandomAccessIterator) types.Size
}
