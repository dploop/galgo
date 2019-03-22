package double_list

import (
	"testing"
)

func TestIterator_Clone(t *testing.T) {
	i := Iterator{p: &Node{}}
	j := i.Clone().(Iterator)
	if &i == &j {
		t.Errorf("&i(%p) == &j(%p)", i, j)
	}
	if i.p != j.p {
		t.Errorf("i.p(%p) != j.p(%p)", i.p, j.p)
	}
}

func BenchmarkIterator_Clone(b *testing.B) {
	i := Iterator{p: &Node{}}
	for n := 0; n < b.N; n++ {
		_ = i.Clone()
	}
}

func TestIterator_ImplClone(t *testing.T) {
	i := Iterator{p: &Node{}}
	j := i.ImplClone()
	if &i == &j {
		t.Errorf("&i(%p) == &j(%p)", i, j)
	}
	if i.p != j.p {
		t.Errorf("i.p(%p) != j.p(%p)", i.p, j.p)
	}
}

func BenchmarkIterator_ImplClone(b *testing.B) {
	i := Iterator{p: &Node{}}
	for n := 0; n < b.N; n++ {
		_ = i.ImplClone()
	}
}

func TestIterator_Read(t *testing.T) {
	i := Iterator{p: &Node{data: 42}}
	data := i.Read().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_Read(b *testing.B) {
	i := Iterator{p: &Node{data: 42}}
	for n := 0; n < b.N; n++ {
		_ = i.Read()
	}
}

func TestIterator_ImplRead(t *testing.T) {
	i := Iterator{p: &Node{data: 42}}
	data := i.ImplRead().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_ImplRead(b *testing.B) {
	i := Iterator{p: &Node{data: 42}}
	for n := 0; n < b.N; n++ {
		_ = i.ImplRead()
	}
}

func TestIterator_Next(t *testing.T) {
	next := &Node{}
	i := Iterator{p: &Node{next: next}}
	i.Next()
	if i.p != next {
		t.Errorf("i.p(%p) != next(%p)", i.p, next)
	}
}

func BenchmarkIterator_Next(b *testing.B) {
	next := &Node{}
	i := Iterator{p: &Node{next: next}}
	i.p.next.next = i.p.next
	for n := 0; n < b.N; n++ {
		i.Next()
	}
}

func TestIterator_Prev(t *testing.T) {
	prev := &Node{}
	i := Iterator{p: &Node{prev: prev}}
	i.Prev()
	if i.p != prev {
		t.Errorf("i.p(%p) != prev(%p)", i.p, prev)
	}
}

func BenchmarkIterator_Prev(b *testing.B) {
	prev := &Node{}
	i := Iterator{p: &Node{prev: prev}}
	i.p.prev.prev = i.p.prev
	for n := 0; n < b.N; n++ {
		i.Prev()
	}
}


func TestIterator_Equal(t *testing.T) {
	i := Iterator{p: &Node{}}
	j := Iterator{p: &Node{}}
	if i.Equal(j) {
		t.Errorf("i(%v) == j(%v)", i, j)
	}
	i.p = j.p
	if !i.Equal(j) {
		t.Errorf("i(%v) != j(%v)", i, j)
	}
}

func BenchmarkIterator_Equal(b *testing.B) {
	i := Iterator{p: &Node{}}
	j := Iterator{p: &Node{}}
	for n := 0; n < b.N; n++ {
		_ = i.Equal(j)
	}
}

func TestIterator_ImplEqual(t *testing.T) {
	i := Iterator{p: &Node{}}
	j := Iterator{p: &Node{}}
	if i.ImplEqual(j) {
		t.Errorf("i(%v) == j(%v)", i, j)
	}
	i.p = j.p
	if !i.ImplEqual(j) {
		t.Errorf("i(%v) != j(%v)", i, j)
	}
}

func BenchmarkIterator_ImplEqual(b *testing.B) {
	i := Iterator{p: &Node{}}
	j := Iterator{p: &Node{}}
	for n := 0; n < b.N; n++ {
		_ = i.ImplEqual(j)
	}
}

func TestIterator_Write(t *testing.T) {
	i := Iterator{p: &Node{}}
	i.Write(42)
	data := i.p.data.(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_Write(b *testing.B) {
	i := Iterator{p: &Node{}}
	for n := 0; n < b.N; n++ {
		i.Write(42)
	}
}

func TestIterator_ImplWrite(t *testing.T) {
	i := Iterator{p: &Node{}}
	i.ImplWrite(42)
	data := i.p.data.(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_ImplWrite(b *testing.B) {
	i := Iterator{p: &Node{}}
	for n := 0; n < b.N; n++ {
		i.ImplWrite(42)
	}
}
