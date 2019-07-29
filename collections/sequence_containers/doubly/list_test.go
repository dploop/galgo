package doubly

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
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("failed to panic")
		}
	}()
	l := NewList()
	l.PushFront(42)
	l.PopFront()
	if !l.Empty() {
		t.Errorf("l(%v) is not empty", l)
	}
	l.PopFront()
}

func BenchmarkList_PopFront(b *testing.B) {
	const loop = 1000000
	var nodes [loop]node
	l := NewList()
	for n := 0; n < b.N; n++ {
		if l.size == 0 {
			b.StopTimer()
			for k := 1; k < loop; k++ {
				nodes[k-1].next = &nodes[k]
				nodes[k].prev = &nodes[k-1]
			}
			l.sent.next = &nodes[0]
			nodes[0].prev = l.sent
			l.sent.prev = &nodes[loop-1]
			nodes[loop-1].next = l.sent
			l.size = loop
			b.StartTimer()
		}
		l.PopFront()
	}
}

func TestList_PushBack(t *testing.T) {
	l := NewList()
	l.PushBack(42)
	data := l.Back().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkList_PushBack(b *testing.B) {
	loop := 1000000
	l := NewList()
	for n := 0; n < b.N; n++ {
		if l.size == loop {
			l.Clear()
		}
		l.PushBack(42)
	}
}

func TestList_Back(t *testing.T) {
	l := NewList()
	l.PushBack(42)
	data := l.Back().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkList_Back(b *testing.B) {
	l := NewList()
	l.PushBack(42)
	for n := 0; n < b.N; n++ {
		_ = l.Back()
	}
}

func TestList_PopBack(t *testing.T) {
	l := NewList()
	l.PushBack(42)
	l.PopBack()
	if !l.Empty() {
		t.Errorf("l(%v) is not empty", l)
	}
}

func BenchmarkList_PopBack(b *testing.B) {
	const loop = 1000000
	var nodes [loop]node
	l := NewList()
	for n := 0; n < b.N; n++ {
		if l.size == 0 {
			b.StopTimer()
			for k := 1; k < loop; k++ {
				nodes[k-1].next = &nodes[k]
				nodes[k].prev = &nodes[k-1]
			}
			nodes[0].prev = l.sent
			l.sent.next = &nodes[0]
			nodes[loop-1].next = l.sent
			l.sent.prev = &nodes[loop-1]
			l.size = loop
			b.StartTimer()
		}
		l.PopBack()
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
	s := l.sent
	begin := l.Begin()
	if begin.n != s {
		t.Errorf("begin.n(%p) != s(%p)", begin.n, s)
	}
	l.PushFront(42)
	begin = l.Begin()
	if begin.n == s {
		t.Errorf("begin.n(%p) == s(%p)", begin.n, s)
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
	s := l.sent
	end := l.End()
	if end.n != s {
		t.Errorf("end.n(%p) != s(%p)", end.n, s)
	}
	l.PushFront(42)
	end = l.End()
	if end.n != s {
		t.Errorf("end.n(%p) != s(%p)", end.n, s)
	}
}

func BenchmarkList_End(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		_ = l.End()
	}
}

func TestList_ReverseBegin(t *testing.T) {
	l := NewList()
	s := l.sent
	rbegin := l.ReverseBegin()
	if rbegin.n != s {
		t.Errorf("rbegin.n(%p) != s(%p)", rbegin.n, s)
	}
	l.PushBack(42)
	rbegin = l.ReverseBegin()
	if rbegin.n == s {
		t.Errorf("rbegin.n(%p) == s(%p)", rbegin.n, s)
	}
}

func BenchmarkList_ReverseBegin(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		_ = l.ReverseBegin()
	}
}

func TestList_ReverseEnd(t *testing.T) {
	l := NewList()
	s := l.sent
	rend := l.ReverseEnd()
	if rend.n != s {
		t.Errorf("rend.n(%p) != s(%p)", rend.n, s)
	}
	l.PushBack(42)
	rend = l.ReverseEnd()
	if rend.n != s {
		t.Errorf("rend.n(%p) != s(%p)", rend.n, s)
	}
}

func BenchmarkList_ReverseEnd(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		_ = l.ReverseEnd()
	}
}

func TestList_Insert(t *testing.T) {
	l := NewList()
	l.PushBack(1)
	i := l.ReverseBegin()
	j := l.Insert(i, 2)
	k := l.Insert(j, 3)
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

func BenchmarkList_Insert(b *testing.B) {
	loop := 1000000
	l := NewList()
	l.PushBack(42)
	i := l.ReverseBegin()
	for n := 0; n < b.N; n++ {
		if l.size == loop {
			l.Clear()
			l.PushBack(42)
			i = l.ReverseBegin()
		}
		_ = l.Insert(i, 42)
	}
}

func TestList_Erase(t *testing.T) {
	l := NewList()
	l.PushBack(3)
	l.PushBack(2)
	l.PushBack(1)
	i := l.ReverseBegin()
	_ = l.Erase(i)
	if l.Size() != 2 {
		t.Errorf("l(%v) is not 2 sized", l)
	}
	_ = l.Erase(i)
	if l.Size() != 1 {
		t.Errorf("l(%v) is not 1 sized", l)
	}
}

func BenchmarkList_Erase(b *testing.B) {
	const loop = 1000000
	var nodes [loop]node
	l := NewList()
	i := l.ReverseBegin()
	for n := 0; n < b.N; n++ {
		if l.size <= 1 {
			b.StopTimer()
			for k := 1; k < loop; k++ {
				nodes[k-1].next = &nodes[k]
				nodes[k].prev = &nodes[k-1]
			}
			nodes[loop-1].next = l.sent
			l.sent.prev = &nodes[loop-1]
			nodes[0].prev = l.sent
			l.sent.next = &nodes[0]
			l.size = loop
			i = l.ReverseBegin()
			b.StartTimer()
		}
		_ = l.Erase(i)
	}
}
