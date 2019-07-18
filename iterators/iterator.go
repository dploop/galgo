package iterators

import (
	"github.com/dploop/gostl/traits"
)

type Iterator interface {
	traits.Cloneable
	traits.Incrementable
}
