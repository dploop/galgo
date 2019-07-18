package vector

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/dploop/gostl/types"
)

func TestNewList(t *testing.T) {
	l := NewList()
	if l == nil {
		t.Errorf("l(%v) == nil", l)
	}
}

func BenchmarkNewList(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		l = NewList()
	}
	_, _ = fmt.Fprint(ioutil.Discard, l)
}

func TestList_Size(t *testing.T) {
	l := NewList()
	if l.Size() != 0 {
		t.Errorf("l(%v) is not 0 sized", l)
	}
	l.PushBack(42)
	if l.Size() != 1 {
		t.Errorf("l(%v) is not 1 sized", l)
	}
}

func BenchmarkList_Size(b *testing.B) {
	l := NewList()
	var tmp int
	for n := 0; n < b.N; n++ {
		if n % 1000000 == 0 {
			l.PushBack(42)
		}
		tmp += l.Size()
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}

func TestList_Empty(t *testing.T) {
	l := NewList()
	if !l.Empty() {
		t.Errorf("l(%v) is not empty", l)
	}
	l.PushBack(42)
	if l.Empty() {
		t.Errorf("l(%v) is empty", l)
	}
}

func BenchmarkList_Empty(b *testing.B) {
	l := NewList()
	var tmp int
	for n := 0; n < b.N; n++ {
		if n % 1000000 == 0 {
			l.PushBack(42)
		}
		if l.Empty() {
			tmp++
		}
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}

func TestList_Get(t *testing.T) {
	l := NewList()
	l.PushBack(42)
	data := l.Get(0).(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkList_Get(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	var tmp int
	for n := 0; n < b.N; n++ {
		l.array[0] = n
		tmp = l.Get(0).(int)
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}

func TestList_Set(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	l.Set(0, 42)
	data := l.Get(0).(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkList_Set(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	for n := 0; n < b.N; n++ {
		l.Set(0, n)
	}
	_, _ = fmt.Fprint(ioutil.Discard, l.Get(0))
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
		if len(l.array) == loop {
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
	var tmp int
	for n := 0; n < b.N; n++ {
		l.array[0] = n
		tmp = l.Back().(int)
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
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
	array := make([]types.Data, loop)
	l := NewList()
	for n := 0; n < b.N; n++ {
		if len(l.array) == 0 {
			b.StopTimer()
			l.array = array
			b.StartTimer()
		}
		l.PopBack()
	}
}

func TestList_Clear(t *testing.T) {
	l := NewList()
	l.PushBack(42)
	l.Clear()
	if !l.Empty() {
		t.Errorf("l(%v) is not empty", l)
	}
}

func BenchmarkList_Clear(b *testing.B) {
	l := NewList()
	for n := 0; n < b.N; n++ {
		l.Clear()
		if n % 1000000 == 0 {
			l.PushBack(0)
		}
	}
	_, _ = fmt.Fprint(ioutil.Discard, l.Size())
}

func TestList_Begin(t *testing.T) {
	l := NewList()
	begin := l.Begin()
	if begin.l != l {
		t.Errorf("begin.l(%p) != l(%p)", begin.l, l)
	}
	l.PushBack(42)
	begin = l.Begin()
	if begin.l != l {
		t.Errorf("begin.l(%p) != l(%p)", begin.l, l)
	}
}

func BenchmarkList_Begin(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	var tmp Iterator
	for n := 0; n < b.N; n++ {
		if n % 1000000 == 0 {
			l.PushBack(0)
		}
		tmp = l.Begin()
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp.n)
}

func TestList_End(t *testing.T) {
	l := NewList()
	end := l.End()
	if end.l != l {
		t.Errorf("end.l(%p) != l(%p)", end.l, l)
	}
	l.PushBack(42)
	end = l.End()
	if end.l != l {
		t.Errorf("end.l(%p) != l(%p)", end.l, l)
	}
}

func BenchmarkList_End(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	var tmp Iterator
	for n := 0; n < b.N; n++ {
		if n % 1000000 == 0 {
			l.PushBack(0)
		}
		tmp = l.End()
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp.n)
}

func TestList_ReverseBegin(t *testing.T) {
	l := NewList()
	rbegin := l.ReverseBegin()
	if rbegin.l != l {
		t.Errorf("rbegin.l(%p) != l(%p)", rbegin.l, l)
	}
	l.PushBack(42)
	rbegin = l.ReverseBegin()
	if rbegin.l != l {
		t.Errorf("rbegin.l(%p) != l(%p)", rbegin.l, l)
	}
}

func BenchmarkList_ReverseBegin(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	var tmp Iterator
	for n := 0; n < b.N; n++ {
		if n % 1000000 == 0 {
			l.PushBack(0)
		}
		tmp = l.ReverseBegin()
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp.n)
}

func TestList_ReverseEnd(t *testing.T) {
	l := NewList()
	rend := l.ReverseEnd()
	if rend.l != l {
		t.Errorf("rend.l(%p) != l(%p)", rend.l, l)
	}
	l.PushBack(42)
	rend = l.ReverseEnd()
	if rend.l != l {
		t.Errorf("rend.l(%p) != l(%p)", rend.l, l)
	}
}

func BenchmarkList_ReverseEnd(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	var tmp Iterator
	for n := 0; n < b.N; n++ {
		if n % 1000000 == 0 {
			l.PushBack(0)
		}
		tmp = l.ReverseEnd()
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp.n)
}
