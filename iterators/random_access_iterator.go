package iterators

import (
	"github.com/dploop/galgo/basics"
)

type RandomAccessIterator interface {
	BidirectionalIterator
	basics.LessThanComparable
	At(diff Diff) Data
	Advance(diff Diff)
	Distance(other RandomAccessIterator) Diff
}
