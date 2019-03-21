package basics

type Swappable interface {
	Swap(other Swappable)
}

func Swap(a Swappable, b Swappable) {
	a.Swap(b)
}
