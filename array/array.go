package array

import (
	"errors"
	"fmt"
	"sync"
)

var (
	// ErrOutOf is returned when the index is out of range.
	ErrOutOf   = errors.New("index out of range")
	ErrNotFind = errors.New("not find")
)

type allow interface {
	int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64 | string | interface{} | *int | *int16 | *int32 | *int64 | *uint | *uint16 | *uint32 | *uint64 | *float32 | *float64 | *string | *interface{}
}

type array[T allow] struct {
	mu   sync.Mutex
	pool []T
}

// New returns a new array.
func New[T allow](cap int) *array[T] {
	return &array[T]{
		pool: make([]T, 0, cap),
	}
}

// Len returns the length of the array.
func (a *array[T]) Len() int {
	a.mu.Lock()
	defer a.mu.Unlock()

	return len(a.pool)
}

// Cap returns the capacity of the array.
func (a *array[T]) Cap() int {
	a.mu.Lock()
	defer a.mu.Unlock()

	return cap(a.pool)
}

// Map applies the function fn to each element of the array.
func (a *array[T]) Map(fn func(T) T) {
	a.mu.Lock()
	defer a.mu.Unlock()

	for i := range a.pool {
		a.pool[i] = fn(a.pool[i])
	}
}

// Filter returns a new array containing all elements for which the function fn returns true.
func (a *array[T]) Filter(fn func(T) bool) []T {
	a.mu.Lock()
	defer a.mu.Unlock()

	var tmp []T
	for i := range a.pool {
		if fn(a.pool[i]) {
			tmp = append(tmp, a.pool[i])
		}
	}
	return tmp
}

// Find returns the first element for which the function fn returns true.
func (a *array[T]) Find(fn func(T) bool) (T, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	for i := range a.pool {
		if fn(a.pool[i]) {
			return a.pool[i], nil
		}
	}
	var zero T
	return zero, ErrNotFind
}

// Each applies the function fn to each element of the array.
func (a *array[T]) Each(fn func(T)) {
	for i := range a.pool {
		fn(a.pool[i])
	}
}

// Push appends a new element to the end of the array.
func (a *array[T]) Push(v T) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.pool = append(a.pool, v)
}

// Unshift prepends a new element to the beginning of the array.
func (a *array[T]) Unshift(v T) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.pool = append([]T{v}, a.pool...)
}

// Shift removes and returns the first element of the array.
func (a *array[T]) Shift() (T, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.pool) == 0 {
		var zero T
		return zero, ErrOutOf
	}

	v := a.pool[0]
	a.pool = a.pool[1:]
	return v, nil
}

// Pop removes and returns the last element of the array.
func (a *array[T]) Pop() (T, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.pool) == 0 {
		var zero T
		return zero, ErrOutOf
	}

	v := a.pool[len(a.pool)-1]
	a.pool = a.pool[:len(a.pool)-1]
	return v, nil
}

// At returns the element at the specified index.
func (a *array[T]) At(i int) (T, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if i < 0 || i >= len(a.pool) {
		var zero T
		return zero, ErrOutOf
	}

	return a.pool[i], nil
}

// String returns a string representation of the array.
func (a *array[T]) String() string {
	a.mu.Lock()
	defer a.mu.Unlock()

	return fmt.Sprintf("%v", a.pool)
}

// Pool returns the underlying pool.
func (a *array[T]) Pool() []T {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.pool
}
