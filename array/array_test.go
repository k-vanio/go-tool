package array_test

import (
	"testing"

	"github.com/k-vanio/go-tool/array"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		a := array.New[int](10)
		assert.NotNil(t, a)
		assert.Equal(t, 0, a.Len())
	})

	t.Run("Each", func(t *testing.T) {
		a := array.New[int](10)
		a.Push(1)
		a.Push(2)
		a.Push(3)
		a.Each(func(v int) {
			assert.Equal(t, v, a.Pool()[v-1])
		})
	})

	t.Run("Map", func(t *testing.T) {
		a := array.New[int](10)
		a.Push(1)
		a.Push(2)
		a.Push(3)
		a.Map(func(v int) int {
			return v * 2
		})
		assert.Equal(t, 2, a.Pool()[0])
		assert.Equal(t, 4, a.Pool()[1])
		assert.Equal(t, 6, a.Pool()[2])
	})

	t.Run("Push", func(t *testing.T) {
		a := array.New[int](10)
		a.Push(1)
		assert.Equal(t, 1, a.Len())
		a.Push(2)
		assert.Equal(t, 2, a.Len())
	})

	t.Run("Filter", func(t *testing.T) {
		a := array.New[int](10)
		a.Push(1)
		a.Push(2)
		a.Push(3)
		tmp := a.Filter(func(v int) bool {
			return v > 1
		})
		assert.Equal(t, 2, tmp[0])
		assert.Equal(t, 3, tmp[1])
	})

	t.Run("Find", func(t *testing.T) {
		a := array.New[int](10)
		a.Push(1)
		a.Push(2)
		a.Push(3)
		v, err := a.Find(func(v int) bool {
			return v > 1
		})
		assert.Equal(t, 2, v)
		assert.Nil(t, err)
		v, err = a.Find(func(v int) bool {
			return v > 3
		})
		assert.Equal(t, 0, v)
		assert.NotNil(t, err)
	})

	t.Run("Unshift", func(t *testing.T) {
		a := array.New[int](10)
		a.Unshift(1)
		assert.Equal(t, 1, a.Len())
		a.Unshift(2)
		assert.Equal(t, 2, a.Len())
	})

	t.Run("Cap", func(t *testing.T) {
		a := array.New[int](2)
		assert.Equal(t, 2, a.Cap())
		a.Push(1)
		assert.Equal(t, 2, a.Cap())
		a.Push(2)
		assert.Equal(t, 2, a.Cap())
	})

	t.Run("Shift", func(t *testing.T) {
		a := array.New[int](10)
		a.Push(1)
		a.Push(2)
		a.Push(3)
		v, err := a.Shift()
		assert.Equal(t, 1, v)
		assert.Nil(t, err)
		v, err = a.Shift()
		assert.Equal(t, 2, v)
		assert.Nil(t, err)
		v, err = a.Shift()
		assert.Equal(t, 3, v)
		assert.Nil(t, err)
		v, err = a.Shift()
		assert.Equal(t, 0, v)
		assert.NotNil(t, err)
	})

	t.Run("Pop", func(t *testing.T) {
		a := array.New[int](10)
		a.Push(1)
		a.Push(2)
		a.Push(3)
		v, err := a.Pop()
		assert.Equal(t, 3, v)
		assert.Nil(t, err)
		v, err = a.Pop()
		assert.Equal(t, 2, v)
		assert.Nil(t, err)
		v, err = a.Pop()
		assert.Equal(t, 1, v)
		assert.Nil(t, err)
		v, err = a.Pop()
		assert.Equal(t, 0, v)
		assert.NotNil(t, err)
	})

	t.Run("At", func(t *testing.T) {
		a := array.New[int](10)
		a.Push(1)
		a.Push(2)
		a.Push(3)
		v, err := a.At(0)
		assert.Equal(t, 1, v)
		assert.Nil(t, err)
		v, err = a.At(1)
		assert.Equal(t, 2, v)
		assert.Nil(t, err)
		v, err = a.At(2)
		assert.Equal(t, 3, v)
		assert.Nil(t, err)
		v, err = a.At(3)
		assert.Equal(t, 0, v)
		assert.NotNil(t, err)
		v, err = a.At(-1)
		assert.Equal(t, 0, v)
		assert.NotNil(t, err)
	})

	t.Run("String", func(t *testing.T) {
		a := array.New[int](10)
		a.Push(1)
		a.Push(2)
		a.Push(3)
		assert.Equal(t, "[1 2 3]", a.String())
	})

	t.Run("Pool", func(t *testing.T) {
		a := array.New[int](10)
		a.Push(1)
		a.Push(2)
		a.Push(3)
		assert.Equal(t, 3, a.Pool()[2])
	})
}
