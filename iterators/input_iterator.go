package iterators

import (
	"github.com/dploop/galgo/basics"
)

type InputIterator interface {
	Iterator
	basics.EqualityComparable
}
