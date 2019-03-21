package iterators

import (
	"github.com/dploop/galgo/basics"
)

type Iterator interface {
	basics.Cloneable
	basics.Swappable
	Readable
	Incrementable
}

func PreIncr(i Iterator) Iterator {
	i.Next()
	return i
}

func PostIncr(i Iterator) Iterator {
	j := i.Clone().(Iterator)
	i.Next()
	return j
}
