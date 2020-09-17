package iterators

import (
	"github.com/dploop/gostl/constraints"
)

type InputIterator interface {
	Iterator
	constraints.EqualityComparable
	constraints.Readable
}
