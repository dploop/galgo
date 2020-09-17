package iterators

import (
	"github.com/dploop/gostl/constraints"
)

type MutableIterator interface {
	constraints.Writeable
}
