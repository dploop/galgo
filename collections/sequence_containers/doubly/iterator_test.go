package doubly

import (
	"testing"
)

func TestIterator_Write(t *testing.T) {
	i := Iterator{n: &node{}}
	i.Write(42)
	data := i.n.data.(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_Write(b *testing.B) {
	i := Iterator{n: &node{}}
	for n := 0; n < b.N; n++ {
		i.Write(42)
	}
}

func TestIterator_Clone(t *testing.T) {
	i := Iterator{n: &node{}}
	j := i.Clone().(Iterator)
	if &i == &j {
		t.Errorf("&i(%p) == &j(%p)", &i, &j)
	}
	if i.n != j.n {
		t.Errorf("i.n(%p) != j.n(%p)", i.n, j.n)
	}
}

func BenchmarkIterator_Clone(b *testing.B) {
	i := Iterator{n: &node{}}
	for n := 0; n < b.N; n++ {
		_ = i.Clone()
	}
}

func TestIterator_ImplClone(t *testing.T) {
	i := Iterator{n: &node{}}
	j := i.ImplClone()
	if &i == &j {
		t.Errorf("&i(%p) == &j(%p)", &i, &j)
	}
	if i.n != j.n {
		t.Errorf("i.n(%p) != j.n(%p)", i.n, j.n)
	}
}

func BenchmarkIterator_ImplClone(b *testing.B) {
	i := Iterator{n: &node{}}
	for n := 0; n < b.N; n++ {
		_ = i.ImplClone()
	}
}

func TestIterator_Next(t *testing.T) {
	next := &node{}
	i := Iterator{n: &node{next: next}}
	i = i.Next().(Iterator)
	if i.n != next {
		t.Errorf("i.n(%p) != next(%p)", i.n, next)
	}
}

func BenchmarkIterator_Next(b *testing.B) {
	i := Iterator{n: &node{}}
	for n := 0; n < b.N; n++ {
		_ = i.Next()
	}
}

func TestIterator_ImplNext(t *testing.T) {
	next := &node{}
	i := Iterator{n: &node{next: next}}
	i = i.ImplNext()
	if i.n != next {
		t.Errorf("i.n(%p) != next(%p)", i.n, next)
	}
}

func BenchmarkIterator_ImplNext(b *testing.B) {
	i := Iterator{n: &node{}}
	for n := 0; n < b.N; n++ {
		_ = i.ImplNext()
	}
}

func TestIterator_Equal(t *testing.T) {
	i := Iterator{n: &node{}}
	j := Iterator{n: &node{}}
	if i.Equal(j) {
		t.Errorf("i(%v) == j(%v)", i, j)
	}
	i.n = j.n
	if !i.Equal(j) {
		t.Errorf("i(%v) != j(%v)", i, j)
	}
}

func BenchmarkIterator_Equal(b *testing.B) {
	i := Iterator{n: &node{}}
	j := Iterator{n: &node{}}
	for n := 0; n < b.N; n++ {
		_ = i.Equal(j)
	}
}

func TestIterator_ImplEqual(t *testing.T) {
	i := Iterator{n: &node{}}
	j := Iterator{n: &node{}}
	if i.ImplEqual(j) {
		t.Errorf("i(%v) == j(%v)", i, j)
	}
	i.n = j.n
	if !i.ImplEqual(j) {
		t.Errorf("i(%v) != j(%v)", i, j)
	}
}

func BenchmarkIterator_ImplEqual(b *testing.B) {
	i := Iterator{n: &node{}}
	j := Iterator{n: &node{}}
	for n := 0; n < b.N; n++ {
		_ = i.ImplEqual(j)
	}
}

func TestIterator_Read(t *testing.T) {
	i := Iterator{n: &node{data: 42}}
	data := i.Read().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_Read(b *testing.B) {
	i := Iterator{n: &node{data: 42}}
	for n := 0; n < b.N; n++ {
		_ = i.Read()
	}
}

func TestIterator_Prev(t *testing.T) {
	prev := &node{}
	i := Iterator{n: &node{prev: prev}}
	i = i.Prev().(Iterator)
	if i.n != prev {
		t.Errorf("i.n(%p) != prev(%p)", i.n, prev)
	}
}

func BenchmarkIterator_Prev(b *testing.B) {
	i := Iterator{n: &node{}}
	for n := 0; n < b.N; n++ {
		_ = i.Prev()
	}
}

func TestIterator_ImplPrev(t *testing.T) {
	prev := &node{}
	i := Iterator{n: &node{prev: prev}}
	i = i.ImplPrev()
	if i.n != prev {
		t.Errorf("i.n(%p) != prev(%p)", i.n, prev)
	}
}

func BenchmarkIterator_ImplPrev(b *testing.B) {
	i := Iterator{n: &node{}}
	for n := 0; n < b.N; n++ {
		_ = i.ImplPrev()
	}
}
