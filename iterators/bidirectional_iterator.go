package iterators

import (
	"github.com/dploop/gostl/constraints"
)

type BidirectionalIterator interface {
	ForwardIterator
	constraints.Decrementable
}
