package single_list

import (
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewList()
	if l == nil {
		t.Errorf("l(%v) == nil", l)
	}
}

func BenchmarkNewList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = NewList()
	}
}

func TestList_Empty(t *testing.T) {
	l := NewList()
	if !l.Empty() {
		t.Errorf("l(%v) is not empty", l)
	}
	l.PushFront(42)
	if l.Empty() {
		t.Errorf("l(%v) is empty", l)
	}
}

func BenchmarkList_Empty(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		_ = l.Empty()
	}
}

func TestList_Size(t *testing.T) {
	l := NewList()
	if l.Size() != 0 {
		t.Errorf("l(%v) is not 0 sized", l)
	}
	l.PushFront(42)
	if l.Size() != 1 {
		t.Errorf("l(%v) is not 1 sized", l)
	}
}

func BenchmarkList_Size(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		_ = l.Size()
	}
}

func TestList_PushFront(t *testing.T) {
	l := NewList()
	l.PushFront(42)
	data := l.Front().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkList_PushFront(b *testing.B) {
	loop := 1000000
	l := NewList()
	for n := 0; n < b.N; n++ {
		if l.size == loop {
			l.Clear()
		}
		l.PushFront(42)
	}
}

func TestList_Front(t *testing.T) {
	l := NewList()
	l.PushFront(42)
	data := l.Front().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkList_Front(b *testing.B) {
	l := NewList()
	l.PushFront(42)
	for n := 0; n < b.N; n++ {
		_ = l.Front()
	}
}

func TestList_PopFront(t *testing.T) {
	l := NewList()
	l.PushFront(42)
	l.PopFront()
	if !l.Empty() {
		t.Errorf("l(%v) is not empty", l)
	}
}

func BenchmarkList_PopFront(b *testing.B) {
	loop := 1000000
	nodes := make([]*Node, loop)
	for k := 0; k < loop; k++ {
		nodes[k] = &Node{}
	}
	l := NewList()
	for n := 0; n < b.N; n++ {
		if l.size == 0 {
			b.StopTimer()
			for k := 1; k < loop; k++ {
				nodes[k - 1].next = nodes[k]
			}
			l.head, l.size = nodes[0], loop
			b.StartTimer()
		}
		l.PopFront()
	}
}

func TestList_Clear(t *testing.T) {
	l := NewList()
	l.PushFront(42)
	l.Clear()
	if !l.Empty() {
		t.Errorf("l(%v) is not empty", l)
	}
}

func BenchmarkList_Clear(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		l.Clear()
	}
}

func TestList_Begin(t *testing.T) {
	l := NewList()
	begin := l.Begin()
	if begin.p != nil {
		t.Errorf("begin.p(%p) != nil", begin.p)
	}
	l.PushFront(42)
	begin = l.Begin()
	if begin.p == nil {
		t.Errorf("begin.p(%p) == nil", begin.p)
	}
}

func BenchmarkList_Begin(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		_ = l.Begin()
	}
}

func TestList_End(t *testing.T) {
	l := NewList()
	end := l.End()
	if end.p != nil {
		t.Errorf("end.p(%p) != nil", end.p)
	}
	l.PushFront(42)
	end = l.End()
	if end.p != nil {
		t.Errorf("end.p(%p) == nil", end.p)
	}
}

func BenchmarkList_End(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		_ = l.End()
	}
}

func TestList_InsertAfter(t *testing.T) {
	l := NewList()
	l.PushFront(1)
	i := l.Begin()
	j := l.InsertAfter(i, 2)
	k := l.InsertAfter(j, 3)
	datai := i.Read().(int)
	if datai != 1 {
		t.Errorf("datai(%v) != 1", datai)
	}
	dataj := j.Read().(int)
	if dataj != 2 {
		t.Errorf("dataj(%v) != 2", dataj)
	}
	datak := k.Read().(int)
	if datak != 3 {
		t.Errorf("datak(%v) != 1", datak)
	}
}

func BenchmarkList_InsertAfter(b *testing.B) {
	loop := 1000000
	l := NewList()
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
	l := NewList()
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
	loop := 1000000
	nodes := make([]*Node, loop)
	for k := 0; k < loop; k++ {
		nodes[k] = &Node{}
	}
	l := NewList()
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		if l.size <= 1 {
			b.StopTimer()
			for k := 1; k < loop; k++ {
				nodes[k - 1].next = nodes[k]
			}
			l.head, l.size = nodes[0], loop
			i = l.Begin()
			b.StartTimer()
		}
		_ = l.EraseAfter(i)
	}
}
