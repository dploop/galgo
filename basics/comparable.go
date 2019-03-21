package basics

type EqualityComparable interface {
	Equal(other EqualityComparable) bool
}

type LessThanComparable interface {
	Less(other LessThanComparable) bool
}

func Equal(a EqualityComparable, b EqualityComparable) bool {
	return a.Equal(b)
}

func NotEqual(a EqualityComparable, b EqualityComparable) bool {
	return !a.Equal(b)
}

func Less(a LessThanComparable, b LessThanComparable) bool {
	return a.Less(b)
}

func NotLess(a LessThanComparable, b LessThanComparable) bool {
	return !a.Less(b)
}

func Greater(a LessThanComparable, b LessThanComparable) bool {
	return b.Less(a)
}

func NotGreater(a LessThanComparable, b LessThanComparable) bool {
	return !b.Less(a)
}
