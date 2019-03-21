package iterators

type BidirectionalIterator interface {
	ForwardIterator
	Decrementable
}

func PreDecr(i BidirectionalIterator) BidirectionalIterator {
	i.Prev()
	return i
}

func PostDecr(i BidirectionalIterator) BidirectionalIterator {
	j := i.Clone().(BidirectionalIterator)
	i.Prev()
	return j
}
