package vector

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestIterator_Write(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	i.Write(42)
	data := i.Read().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_Write(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		i.Write(n)
	}
	_, _ = fmt.Fprint(ioutil.Discard, i.Read())
}

func TestIterator_Clone(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := i.Clone().(Iterator)
	if &i == &j {
		t.Errorf("&i(%p) == &j(%p)", &i, &j)
	}
	if i.l != j.l {
		t.Errorf("i.l(%p) != j.l(%p)", i.l, j.l)
	}
	if i.n != j.n {
		t.Errorf("i.n(%v) != j.n(%v)", i.n, j.n)
	}
}

func BenchmarkIterator_Clone(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		i = i.Clone().(Iterator)
		i.n++
	}
	_, _ = fmt.Fprint(ioutil.Discard, i.n)
}

func TestIterator_ImplClone(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := i.ImplClone()
	if &i == &j {
		t.Errorf("&i(%p) == &j(%p)", &i, &j)
	}
	if i.l != j.l {
		t.Errorf("i.l(%p) != j.l(%p)", i.l, j.l)
	}
	if i.n != j.n {
		t.Errorf("i.n(%v) != j.n(%v)", i.n, j.n)
	}
}

func BenchmarkIterator_ImplClone(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		i = i.ImplClone()
		i.n++
	}
	_, _ = fmt.Fprint(ioutil.Discard, i.n)
}

func TestIterator_Next(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	i = i.Next().(Iterator)
	if i.n != 1 {
		t.Errorf("i.n(%v) != 1", i.n)
	}
}

func BenchmarkIterator_Next(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		i = i.Next().(Iterator)
	}
	_, _ = fmt.Fprint(ioutil.Discard, i.n)
}

func TestIterator_ImplNext(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	i = i.ImplNext()
	if i.n != 1 {
		t.Errorf("i.n(%v) != 1", i.n)
	}
}

func BenchmarkIterator_ImplNext(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		i = i.ImplNext()
	}
	_, _ = fmt.Fprint(ioutil.Discard, i.n)
}

func TestIterator_Equal(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.Begin()
	if !i.Equal(j) {
		t.Errorf("i(%v) != j(%v)", i, j)
	}
}

func BenchmarkIterator_Equal(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.Begin()
	var tmp int
	for n := 0; n < b.N; n++ {
		if i.Equal(j) {
			tmp++
		}
		i.n = (i.n + n) % 2
		j.n = (j.n + n) % 3
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}

func TestIterator_ImplEqual(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.Begin()
	if !i.ImplEqual(j) {
		t.Errorf("i(%v) != j(%v)", i, j)
	}
}

func BenchmarkIterator_ImplEqual(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.Begin()
	var tmp int
	for n := 0; n < b.N; n++ {
		if i.Equal(j) {
			tmp++
		}
		i.n = (i.n + n) % 2
		j.n = (j.n + n) % 3
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}

func TestIterator_Read(t *testing.T) {
	l := NewList()
	l.PushBack(42)
	i := l.Begin()
	data := i.Read().(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_Read(b *testing.B) {
	l := NewList()
	l.PushBack(1)
	i := l.Begin()
	var tmp int
	for n := 0; n < b.N; n++ {
		l.array[0] = n
		tmp = i.Read().(int)
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}

func TestIterator_Prev(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	i = i.Prev().(Iterator)
	if i.n != -1 {
		t.Errorf("i.n(%v) != -1", i.n)
	}
}

func BenchmarkIterator_Prev(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		i = i.Prev().(Iterator)
	}
	_, _ = fmt.Fprint(ioutil.Discard, i.n)
}

func TestIterator_ImplPrev(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	i = i.ImplPrev()
	if i.n != -1 {
		t.Errorf("i.n(%v) != -1", i.n)
	}
}

func BenchmarkIterator_ImplPrev(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		i = i.ImplPrev()
	}
	_, _ = fmt.Fprint(ioutil.Discard, i.n)
}

func TestIterator_Less(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.End()
	if !i.Less(j) {
		t.Errorf("i(%v) should be less than j(%v)", i, j)
	}
}

func BenchmarkIterator_Less(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.End()
	var tmp int
	for n := 0; n < b.N; n++ {
		if i.Less(j) {
			tmp++
		}
		i.n = (i.n + n) % 2
		j.n = (j.n + n) % 3
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}

func TestIterator_ImplLess(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.End()
	if !i.ImplLess(j) {
		t.Errorf("i(%v) should be less than j(%v)", i, j)
	}
}

func BenchmarkIterator_ImplLess(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.End()
	var tmp int
	for n := 0; n < b.N; n++ {
		if i.ImplLess(j) {
			tmp++
		}
		i.n = (i.n + n) % 2
		j.n = (j.n + n) % 3
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}

func TestIterator_At(t *testing.T) {
	l := NewList()
	l.PushBack(42)
	i := l.Begin()
	data := i.At(0).(int)
	if data != 42 {
		t.Errorf("data(%v) != 42", data)
	}
}

func BenchmarkIterator_At(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	var tmp int
	for n := 0; n < b.N; n++ {
		l.array[0] = n
		tmp = i.At(0).(int)
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}

func TestIterator_Advance(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	i = i.Advance(1).(Iterator)
	if i.n != 1 {
		t.Errorf("i.n(%v) != 1", i.n)
	}
}

func BenchmarkIterator_Advance(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		i = i.Advance(1).(Iterator)
	}
	_, _ = fmt.Fprint(ioutil.Discard, i.n)
}

func TestIterator_ImplAdvance(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	i = i.ImplAdvance(1)
	if i.n != 1 {
		t.Errorf("i.n(%v) != 1", i.n)
	}
}

func BenchmarkIterator_ImplAdvance(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	for n := 0; n < b.N; n++ {
		i = i.ImplAdvance(1)
	}
	_, _ = fmt.Fprint(ioutil.Discard, i.n)
}

func TestIterator_Distance(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.End()
	dist := i.Distance(j)
	if dist != 1 {
		t.Errorf("dist(%v) != 1", dist)
	}
}

func BenchmarkIterator_Distance(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.End()
	var tmp int
	for n := 0; n < b.N; n++ {
		tmp += i.Distance(j)
		i.n = (i.n + n) % 2
		j.n = (j.n + n) % 3
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}

func TestIterator_ImplDistance(t *testing.T) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.End()
	dist := i.ImplDistance(j)
	if dist != 1 {
		t.Errorf("dist(%v) != 1", dist)
	}
}

func BenchmarkIterator_ImplDistance(b *testing.B) {
	l := NewList()
	l.PushBack(0)
	i := l.Begin()
	j := l.End()
	var tmp int
	for n := 0; n < b.N; n++ {
		tmp += i.ImplDistance(j)
		i.n = (i.n + n) % 2
		j.n = (j.n + n) % 3
	}
	_, _ = fmt.Fprint(ioutil.Discard, tmp)
}
