package iterators

import (
	"github.com/dploop/gostl/constraints"
)

type Iterator interface {
	constraints.Cloneable
	constraints.Incrementable
}
