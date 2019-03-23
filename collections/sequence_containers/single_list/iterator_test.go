package single_list

import (
	"testing"
)

func TestIterator_Clone(t *testing.T) {
	i := Iterator{&node{}}
	j := i.Clone().(Iterator)
	if &i == &j {
		t.Errorf("&i(%p) == &j(%p)", i, j)
	}
	if i.node != j.node {
		t.Errorf("i.node(%p) != j.node(%p)", i.node, j.node)
	}
}

func BenchmarkIterator_Clone(b *testing.B) {
	i := Iterator{&node{}}
	for n := 0; n < b.N; n++ {
		_ = i.Clone()
	}
}

func TestIterator_ImplClone(t *testing.T) {
	i := Iterator{&node{}}
	j := i.ImplClone()
	if &i == &j {
		t.Errorf("&i(%p) == &j(%p)", i, j)
	}
	if i.node != j.node {
		t.Errorf("i.node(%p) != j.node(%p)", i.node, j.node)
	}
}

func BenchmarkIterator_ImplClone(b *testing.B) {
	i := Iterator{&node{}}
	for n := 0; n < b.N; n++ {
		_ = i.ImplClone()
	}
}

func TestIterator_Read(t *testing.T) {
	i := Iterator{&node{data: 42}}
	data := i.Read().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_Read(b *testing.B) {
	i := Iterator{&node{data: 42}}
	for n := 0; n < b.N; n++ {
		_ = i.Read()
	}
}

func TestIterator_ImplRead(t *testing.T) {
	i := Iterator{&node{data: 42}}
	data := i.ImplRead().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_ImplRead(b *testing.B) {
	i := Iterator{&node{data: 42}}
	for n := 0; n < b.N; n++ {
		_ = i.ImplRead()
	}
}

func TestIterator_Next(t *testing.T) {
	next := &node{}
	i := Iterator{&node{next: next}}
	i.Next()
	if i.node != next {
		t.Errorf("i.node(%p) != next(%p)", i.node, next)
	}
}

func BenchmarkIterator_Next(b *testing.B) {
	next := &node{}
	i := Iterator{&node{next: next}}
	i.next.next = i.next
	for n := 0; n < b.N; n++ {
		i.Next()
	}
}

func TestIterator_Equal(t *testing.T) {
	i := Iterator{&node{}}
	j := Iterator{&node{}}
	if i.Equal(j) {
		t.Errorf("i(%v) == j(%v)", i, j)
	}
	i.node = j.node
	if !i.Equal(j) {
		t.Errorf("i(%v) != j(%v)", i, j)
	}
}

func BenchmarkIterator_Equal(b *testing.B) {
	i := Iterator{&node{}}
	j := Iterator{&node{}}
	for n := 0; n < b.N; n++ {
		_ = i.Equal(j)
	}
}

func TestIterator_ImplEqual(t *testing.T) {
	i := Iterator{&node{}}
	j := Iterator{&node{}}
	if i.ImplEqual(j) {
		t.Errorf("i(%v) == j(%v)", i, j)
	}
	i.node = j.node
	if !i.ImplEqual(j) {
		t.Errorf("i(%v) != j(%v)", i, j)
	}
}

func BenchmarkIterator_ImplEqual(b *testing.B) {
	i := Iterator{&node{}}
	j := Iterator{&node{}}
	for n := 0; n < b.N; n++ {
		_ = i.ImplEqual(j)
	}
}

func TestIterator_Write(t *testing.T) {
	i := Iterator{&node{}}
	i.Write(42)
	data := i.data.(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_Write(b *testing.B) {
	i := Iterator{&node{}}
	for n := 0; n < b.N; n++ {
		i.Write(42)
	}
}

func TestIterator_ImplWrite(t *testing.T) {
	i := Iterator{&node{}}
	i.ImplWrite(42)
	data := i.data.(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_ImplWrite(b *testing.B) {
	i := Iterator{&node{}}
	for n := 0; n < b.N; n++ {
		i.ImplWrite(42)
	}
}
