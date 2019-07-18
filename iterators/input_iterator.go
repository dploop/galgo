package iterators

import (
	"github.com/dploop/gostl/traits"
)

type InputIterator interface {
	Iterator
	traits.EqualityComparable
	traits.Readable
}
