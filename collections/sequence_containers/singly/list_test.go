package singly

import (
	"testing"
)

func TestNewList(t *testing.T) {
	l := New()
	if l == nil {
		t.Errorf("l(%v) == nil", l)
	}
}

func BenchmarkNewList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = New()
	}
}

func TestList_Size(t *testing.T) {
	l := New()
	if l.Size() != 0 {
		t.Errorf("l(%v) is not 0 sized", l)
	}
	l.PushFront(42)
	if l.Size() != 1 {
		t.Errorf("l(%v) is not 1 sized", l)
	}
}

func BenchmarkList_Size(b *testing.B) {
	l := New()
	for n := 0; n < b.N; n++ {
		_ = l.Size()
	}
}

func TestList_Empty(t *testing.T) {
	l := New()
	if !l.Empty() {
		t.Errorf("l(%v) is not empty", l)
	}
	l.PushFront(42)
	if l.Empty() {
		t.Errorf("l(%v) is empty", l)
	}
}

func BenchmarkList_Empty(b *testing.B) {
	l := New()
	for n := 0; n < b.N; n++ {
		_ = l.Empty()
	}
}

func TestList_PushFront(t *testing.T) {
	l := New()
	l.PushFront(42)
	data := l.Front().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkList_PushFront(b *testing.B) {
	loop := 1000000
	l := New()
	for n := 0; n < b.N; n++ {
		if l.size == loop {
			l.Clear()
		}
		l.PushFront(42)
	}
}

func TestList_Front(t *testing.T) {
	l := New()
	l.PushFront(42)
	data := l.Front().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkList_Front(b *testing.B) {
	l := New()
	l.PushFront(42)
	for n := 0; n < b.N; n++ {
		_ = l.Front()
	}
}

func TestList_PopFront(t *testing.T) {
	l := New()
	l.PushFront(42)
	l.PopFront()
	if !l.Empty() {
		t.Errorf("l(%v) is not empty", l)
	}
}

func BenchmarkList_PopFront(b *testing.B) {
	loop := 1000000
	nodes := make([]*node, loop)
	for k := 0; k < loop; k++ {
		nodes[k] = &node{}
	}
	l := New()
	for n := 0; n < b.N; n++ {
		if l.size == 0 {
			b.StopTimer()
			for k := 1; k < loop; k++ {
				nodes[k-1].next = nodes[k]
			}
			l.sentinel.next = nodes[0]
			l.size = loop
			b.StartTimer()
		}
		l.PopFront()
	}
}

func TestList_Clear(t *testing.T) {
	l := New()
	l.PushFront(42)
	l.Clear()
	if !l.Empty() {
		t.Errorf("l(%v) is not empty", l)
	}
}

func BenchmarkList_Clear(b *testing.B) {
	l := New()
	for n := 0; n < b.N; n++ {
		l.Clear()
	}
}

func TestList_Begin(t *testing.T) {
	l := New()
	begin := l.Begin()
	if begin.n != nil {
		t.Errorf("begin.n(%p) != nil", begin.n)
	}
	l.PushFront(42)
	begin = l.Begin()
	if begin.n == nil {
		t.Errorf("begin.n(%p) == nil", begin.n)
	}
}

func BenchmarkList_Begin(b *testing.B) {
	l := New()
	for n := 0; n < b.N; n++ {
		_ = l.Begin()
	}
}

func TestList_End(t *testing.T) {
	l := New()
	end := l.End()
	if end.n != nil {
		t.Errorf("end.n(%p) != nil", end.n)
	}
	l.PushFront(42)
	end = l.End()
	if end.n != nil {
		t.Errorf("end.n(%p) != nil", end.n)
	}
}

func BenchmarkList_End(b *testing.B) {
	l := New()
	for n := 0; n < b.N; n++ {
		_ = l.End()
	}
}

func TestList_InsertAfter(t *testing.T) {
	l := New()
	l.PushFront(1)
	i := l.Begin()
	j := l.InsertAfter(i, 2)
	k := l.InsertAfter(j, 3)
	di := i.Read().(int)
	if di != 1 {
		t.Errorf("di(%v) != 1", di)
	}
	dj := j.Read().(int)
	if dj != 2 {
		t.Errorf("dj(%v) != 2", dj)
	}
	dk := k.Read().(int)
	if dk != 3 {
		t.Errorf("dk(%v) != 1", dk)
	}
}

func BenchmarkList_InsertAfter(b *testing.B) {
	loop := 1000000
	l := New()
	l.PushFront(42)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		if l.size == loop {
			l.Clear()
			l.PushFront(42)
			i = l.Begin()
		}
		_ = l.InsertAfter(i, 42)
	}
}

func TestList_EraseAfter(t *testing.T) {
	l := New()
	l.PushFront(3)
	l.PushFront(2)
	l.PushFront(1)
	i := l.Begin()
	_ = l.EraseAfter(i)
	if l.Size() != 2 {
		t.Errorf("l(%v) is not 2 sized", l)
	}
	_ = l.EraseAfter(i)
	if l.Size() != 1 {
		t.Errorf("l(%v) is not 1 sized", l)
	}
}

func BenchmarkList_EraseAfter(b *testing.B) {
	const loop = 1000000
	var nodes [loop]node
	l := New()
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		if l.size <= 1 {
			b.StopTimer()
			for k := 1; k < loop; k++ {
				nodes[k-1].next = &nodes[k]
			}
			l.sentinel.next = &nodes[0]
			l.size = loop
			i = l.Begin()
			b.StartTimer()
		}
		_ = l.EraseAfter(i)
	}
}
