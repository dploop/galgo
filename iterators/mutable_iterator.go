package iterators

import (
	"github.com/dploop/gostl/traits"
)

type MutableIterator interface {
	traits.Writeable
}
