package iterators

import (
	"github.com/dploop/gostl/traits"
	"github.com/dploop/gostl/types"
)

type RandomAccessIterator interface {
	BidirectionalIterator
	traits.LessThanComparable
	At(diff types.Size) types.Data
	Advance(diff types.Size) RandomAccessIterator
	Distance(other RandomAccessIterator) types.Size
}
