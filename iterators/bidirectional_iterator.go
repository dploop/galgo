package iterators

import (
	"github.com/dploop/gostl/traits"
)

type BidirectionalIterator interface {
	ForwardIterator
	traits.Decrementable
}
