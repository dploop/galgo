package basics

type Cloneable interface {
	Clone() Cloneable
}

func Clone(c Cloneable) Cloneable {
	return c.Clone()
}
