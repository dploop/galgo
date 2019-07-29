package avl

import (
	"github.com/dploop/gostl/traits"
	"github.com/dploop/gostl/types"
)

type Tree struct {
	sent *node
	size types.Size
	comp traits.LessThan
}

const (
	leftHeavy  = -1
	balanced   = 0
	rightHeavy = +1
)

type node struct {
	parent *node
	left   *node
	right  *node
	factor int8
	data   types.Data
}

func New(comp traits.LessThan) *Tree {
	return &Tree{
		sent: &node{},
		comp: comp,
	}
}

func (t *Tree) Size() types.Size {
	return t.size
}

func (t *Tree) Empty() bool {
	return t.Size() == 0
}

func (t *Tree) Begin() Iterator {
	return Iterator{n: minimum(t.sent)}
}

func (t *Tree) End() Iterator {
	return Iterator{n: t.sent}
}

func (t *Tree) ReverseBegin() Iterator {
	return Iterator{n: maximum(t.sent)}
}

func (t *Tree) ReverseEnd() Iterator {
	return Iterator{n: t.sent}
}

func (t *Tree) CountUnique(v types.Data) types.Size {
	lb := t.LowerBound(v)
	if lb != t.End() && !t.comp(v, lb.Read()) {
		return 1
	}
	return 0
}

func (t *Tree) CountMulti(v types.Data) (c types.Size) {
	lb, ub := t.LowerBound(v), t.UpperBound(v)
	for lb != ub {
		c++
		lb = lb.ImplNext()
	}
	return c
}

func (t *Tree) Find(v types.Data) Iterator {
	lb := t.LowerBound(v)
	if lb != t.End() && !t.comp(v, lb.Read()) {
		return lb
	}
	return t.End()
}

func (t *Tree) Contains(v types.Data) bool {
	lb := t.LowerBound(v)
	if lb != t.End() && !t.comp(v, lb.Read()) {
		return true
	}
	return false
}

func (t *Tree) EqualRangeUnique(v types.Data) (Iterator, Iterator) {
	lb := t.LowerBound(v)
	if lb != t.End() && !t.comp(v, lb.Read()) {
		return lb, lb.ImplNext()
	}
	return lb, lb
}

func (t *Tree) EqualRangeMulti(v types.Data) (Iterator, Iterator) {
	return t.LowerBound(v), t.UpperBound(v)
}

func (t *Tree) LowerBound(v types.Data) Iterator {
	return Iterator{n: t.lowerBound(v)}
}

func (t *Tree) UpperBound(v types.Data) Iterator {
	return Iterator{n: t.upperBound(v)}
}

func (t *Tree) Clear() {
	t.sent.left = nil
	t.size = 0
}

func (t *Tree) InsertUnique(v types.Data) (Iterator, bool) {
	lb := t.LowerBound(v)
	if lb != t.End() && !t.comp(v, lb.Read()) {
		return t.End(), false
	}
	z := &node{data: v}
	t.insert(z)
	return Iterator{n: z}, true
}

func (t *Tree) InsertMulti(v types.Data) Iterator {
	z := &node{data: v}
	t.insert(z)
	return Iterator{n: z}
}

func (t *Tree) Delete(i Iterator) Iterator {
	r := i.ImplNext()
	t.delete(i.n)
	return r
}

func (t *Tree) lowerBound(v types.Data) *node {
	lb := t.sent
	for x := lb.left; x != nil; {
		if !t.comp(x.data, v) {
			lb = x
			x = x.left
		} else {
			x = x.right
		}
	}
	return lb
}

func (t *Tree) upperBound(v types.Data) *node {
	ub := t.sent
	for x := ub.left; x != nil; {
		if t.comp(v, x.data) {
			ub = x
			x = x.left
		} else {
			x = x.right
		}
	}
	return ub
}

func (t *Tree) insert(z *node) {
	x := t.sent
	less := true
	y := x.left
	for y != nil {
		x = y
		less = t.comp(z.data, y.data)
		if less {
			y = y.left
		} else {
			y = y.right
		}
	}
	z.parent = x
	if less {
		x.left = z
	} else {
		x.right = z
	}
	t.balanceAfterInsert(x, z)
	t.size++
}

func (t *Tree) balanceAfterInsert(x *node, z *node) {
	var g, n *node
	for ; x != t.sent; x = g {
		g = x.parent
		if z == x.right {
			switch x.factor {
			case balanced:
				x.factor, z = rightHeavy, x
				continue
			case rightHeavy:
				if z.factor == leftHeavy {
					n = rotateRightLeft(x, z)
				} else {
					n = rotateLeft(x, z)
				}
			default:
				x.factor = balanced
				return
			}
		} else {
			switch x.factor {
			case balanced:
				x.factor, z = leftHeavy, x
				continue
			case leftHeavy:
				if z.factor == rightHeavy {
					n = rotateLeftRight(x, z)
				} else {
					n = rotateRight(x, z)
				}
			default:
				x.factor = balanced
				return
			}
		}
		n.parent = g
		if x == g.left {
			g.left = n
		} else {
			g.right = n
		}
		return
	}
}

func (t *Tree) delete(z *node) {
	var x, n *node
	if z.right == nil {
		x, n = z.parent, z.left
		transplant(z, n)
	} else {
		y := minimum(z.right)
		x, n = y, y.right
		if y.parent != z {
			x = y.parent
			transplant(y, n)
			y.right = z.right
			y.right.parent = y
		}
		transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.factor = z.factor
	}
	t.balanceAfterDelete(x, n)
	t.size--
}

func (t *Tree) balanceAfterDelete(x *node, n *node) {
	var g, z *node
	var b int8
	for ; x != t.sent; x = g {
		g = x.parent
		if n == x.left {
			switch x.factor {
			case leftHeavy:
				x.factor, n = balanced, x
				continue
			case rightHeavy:
				z = x.right
				if b = z.factor; b == leftHeavy {
					n = rotateRightLeft(x, z)
				} else {
					n = rotateLeft(x, z)
				}
			default:
				x.factor = rightHeavy
				return
			}
		} else {
			switch x.factor {
			case rightHeavy:
				x.factor, n = balanced, x
				continue
			case leftHeavy:
				z = x.left
				if b = z.factor; b == rightHeavy {
					n = rotateLeftRight(x, z)
				} else {
					n = rotateRight(x, z)
				}
			default:
				x.factor = leftHeavy
				return
			}
		}
		n.parent = g
		if x == g.left {
			g.left = n
		} else {
			g.right = n
		}
		if b == balanced {
			return
		}
	}
}

func transplant(u *node, v *node) {
	if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func minimum(x *node) *node {
	for x.left != nil {
		x = x.left
	}
	return x
}

func maximum(x *node) *node {
	for x.right != nil {
		x = x.right
	}
	return x
}

func successor(x *node) *node {
	if x.right != nil {
		return minimum(x.right)
	}
	for x == x.parent.right {
		x = x.parent
	}
	return x.parent
}

func predecessor(x *node) *node {
	if x.left != nil {
		return maximum(x.left)
	}
	for x == x.parent.left {
		x = x.parent
	}
	return x.parent
}

func rotateLeft(x *node, z *node) *node {
	t23 := z.left
	x.right = t23
	if t23 != nil {
		t23.parent = x
	}
	z.left = x
	x.parent = z
	if z.factor == balanced {
		x.factor, z.factor = rightHeavy, leftHeavy
	} else {
		x.factor, z.factor = balanced, balanced
	}
	return z
}

func rotateRight(x *node, z *node) *node {
	t23 := z.right
	x.left = t23
	if t23 != nil {
		t23.parent = x
	}
	z.right = x
	x.parent = z
	if z.factor == balanced {
		x.factor, z.factor = leftHeavy, rightHeavy
	} else {
		x.factor, z.factor = balanced, balanced
	}
	return z
}

func rotateRightLeft(x *node, z *node) *node {
	y := z.left
	t3 := y.right
	if t3 != nil {
		t3.parent = z
	}
	y.right = z
	z.parent = y
	t2 := y.left
	if t2 != nil {
		t2.parent = x
	}
	y.left = x
	x.parent = y
	switch y.factor {
	case rightHeavy:
		x.factor, z.factor = leftHeavy, balanced
	case leftHeavy:
		x.factor, z.factor = balanced, rightHeavy
	default:
		x.factor, z.factor = balanced, balanced
	}
	y.factor = balanced
	return y
}

func rotateLeftRight(x *node, z *node) *node {
	y := z.right
	t3 := y.left
	if t3 != nil {
		t3.parent = z
	}
	y.left = z
	z.parent = y
	t2 := y.right
	if t2 != nil {
		t2.parent = x
	}
	y.right = x
	x.parent = y
	switch y.factor {
	case leftHeavy:
		x.factor, z.factor = rightHeavy, balanced
	case rightHeavy:
		x.factor, z.factor = balanced, leftHeavy
	default:
		x.factor, z.factor = balanced, balanced
	}
	y.factor = balanced
	return y
}
