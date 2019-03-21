package iterators

type BidirectionalIterator interface {
	ForwardIterator
	Decrementable
}
