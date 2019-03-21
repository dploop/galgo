package iterators

import (
	"github.com/dploop/galgo/basics"
)

type RandomAccessIterator interface {
	BidirectionalIterator
	basics.LessThanComparable
	At(diff Size) Data
	Advance(diff Size)
	Distance(other RandomAccessIterator) Size
}
