package initial_test

import (
	"testing"

	"github.com/dploop/gostl/collections/sequence_containers/initial"
	"github.com/stretchr/testify/assert"
)

func TestIterator_Write(t *testing.T) {
	i := newIterator()
	assert.Equal(t, 0, i.Read())
	i.Write(1)
	assert.Equal(t, 1, i.Read())
}

func BenchmarkIterator_Write(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		i.Write(1)
	}
}

func TestIterator_Clone(t *testing.T) {
	i := newIterator()
	j := i.Clone().(initial.Iterator)
	assert.True(t, &i != &j)
	assert.Equal(t, 0, i.Read())
	assert.Equal(t, 0, j.Read())
	j.Write(1)
	assert.Equal(t, 1, i.Read())
	assert.Equal(t, 1, j.Read())
}

func BenchmarkIterator_Clone(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Clone()
	}
}

func TestIterator_ImplClone(t *testing.T) {
	i := newIterator()
	j := i.ImplClone()
	assert.True(t, &i != &j)
	assert.Equal(t, 0, i.Read())
	assert.Equal(t, 0, j.Read())
	j.Write(1)
	assert.Equal(t, 1, i.Read())
	assert.Equal(t, 1, j.Read())
}

func BenchmarkIterator_ImplClone(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.ImplClone()
	}
}

func TestIterator_Next(t *testing.T) {
	l := newList(1, 2)
	i := l.Begin()
	assert.Equal(t, 1, i.Read())
	i = i.Next().(initial.Iterator)
	assert.Equal(t, 2, i.Read())
	i = i.Next().(initial.Iterator)
	assert.Equal(t, l.End(), i)
}

func BenchmarkIterator_Next(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Next()
	}
}

func TestIterator_ImplNext(t *testing.T) {
	l := newList(1, 2)
	i := l.Begin()
	assert.Equal(t, 1, i.Read())
	i = i.ImplNext()
	assert.Equal(t, 2, i.Read())
	i = i.ImplNext()
	assert.Equal(t, l.End(), i)
}

func BenchmarkIterator_ImplNext(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.ImplNext()
	}
}

func TestIterator_Equal(t *testing.T) {
	i := newIterator()
	assert.True(t, i.Equal(i))
	j := newIterator()
	assert.False(t, i.Equal(j))
}

func BenchmarkIterator_Equal(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Equal(i)
	}
}

func TestIterator_ImplEqual(t *testing.T) {
	i := newIterator()
	assert.True(t, i.ImplEqual(i))
	j := newIterator()
	assert.False(t, i.ImplEqual(j))
}

func BenchmarkIterator_ImplEqual(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.ImplEqual(i)
	}
}

func TestIterator_Read(t *testing.T) {
	i := newIterator()
	i.Write(1)
	assert.Equal(t, 1, i.Read())
}

func BenchmarkIterator_Read(b *testing.B) {
	i := newIterator()
	for n := 0; n < b.N; n++ {
		_ = i.Read()
	}
}
