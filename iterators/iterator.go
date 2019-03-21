package iterators

import (
	"github.com/dploop/galgo/basics"
)

type Iterator interface {
	basics.Cloneable
	Readable
	Incrementable
}
