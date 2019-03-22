package double_list

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
				nodes[k].prev = nodes[k - 1]
			}
			l.sent.next = nodes[0]
			nodes[0].prev = &l.sent
			l.sent.prev = nodes[loop - 1]
			nodes[loop - 1].next = &l.sent
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
				nodes[k].prev = nodes[k - 1]
			}
			nodes[0].prev = &l.sent
			l.sent.next = nodes[0]
			nodes[loop - 1].next = &l.sent
			l.sent.prev = nodes[loop - 1]
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
	sent := &l.sent
	begin := l.Begin()
	if begin.p != sent {
		t.Errorf("begin.p(%p) != sent(%p)", begin.p, sent)
	}
	l.PushFront(42)
	begin = l.Begin()
	if begin.p == sent {
		t.Errorf("begin.p(%p) == sent(%p)", begin.p, sent)
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
	sent := &l.sent
	end := l.End()
	if end.p != sent {
		t.Errorf("end.p(%p) != sent(%p)", end.p, sent)
	}
	l.PushFront(42)
	end = l.End()
	if end.p != sent {
		t.Errorf("end.p(%p) != sent(%p)", end.p, sent)
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
	sent := &l.sent
	rbegin := l.ReverseBegin()
	if rbegin.p != sent {
		t.Errorf("rbegin.p(%p) != sent(%p)", rbegin.p, sent)
	}
	l.PushBack(42)
	rbegin = l.ReverseBegin()
	if rbegin.p == sent {
		t.Errorf("rbegin.p(%p) == sent(%p)", rbegin.p, sent)
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
	sent := &l.sent
	rend := l.ReverseEnd()
	if rend.p != sent {
		t.Errorf("rend.p(%p) != sent(%p)", rend.p, sent)
	}
	l.PushBack(42)
	rend = l.ReverseEnd()
	if rend.p != sent {
		t.Errorf("rend.p(%p) != sent(%p)", rend.p, sent)
	}
}

func BenchmarkList_ReverseEnd(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		_ = l.ReverseEnd()
	}
}

func TestList_InsertBefore(t *testing.T) {
	l := NewList()
	l.PushBack(1)
	i := l.ReverseBegin()
	j := l.InsertBefore(i, 2)
	k := l.InsertBefore(j, 3)
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

func BenchmarkList_InsertBefore(b *testing.B) {
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
		_ = l.InsertBefore(i, 42)
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

func TestList_EraseBefore(t *testing.T) {
	l := NewList()
	l.PushBack(3)
	l.PushBack(2)
	l.PushBack(1)
	i := l.ReverseBegin()
	_ = l.EraseBefore(i)
	if l.Size() != 2 {
		t.Errorf("l(%v) is not 2 sized", l)
	}
	_ = l.EraseBefore(i)
	if l.Size() != 1 {
		t.Errorf("l(%v) is not 1 sized", l)
	}
}

func BenchmarkList_EraseBefore(b *testing.B) {
	loop := 1000000
	nodes := make([]*Node, loop)
	for k := 0; k < loop; k++ {
		nodes[k] = &Node{}
	}
	l := NewList()
	i := l.ReverseBegin()
	for n := 0; n < b.N; n++ {
		if l.size <= 1 {
			b.StopTimer()
			for k := 1; k < loop; k++ {
				nodes[k - 1].next = nodes[k]
				nodes[k].prev = nodes[k - 1]
			}
			nodes[loop - 1].next = &l.sent
			l.sent.prev = nodes[loop - 1]
			nodes[0].prev = &l.sent
			l.sent.next = nodes[0]
			l.size = loop
			i = l.ReverseBegin()
			b.StartTimer()
		}
		_ = l.EraseBefore(i)
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
				nodes[k].prev = nodes[k - 1]
			}
			nodes[loop - 1].next = &l.sent
			l.sent.prev = nodes[loop - 1]
			nodes[0].prev = &l.sent
			l.sent.next = nodes[0]
			l.size = loop
			i = l.Begin()
			b.StartTimer()
		}
		_ = l.EraseAfter(i)
	}
}
