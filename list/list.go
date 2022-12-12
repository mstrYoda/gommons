package list

import "fmt"

// List[T] is an interface for safe statically typing.
type List[T any] interface {
	At(int) T
	Set(int, T)
	Len() int
	Cap() int
	Empty() bool
	String() string
	Clear()
	Append(...T)
	Clone() List[T]
}

// list[T] is an wrapper for the Go slices.
type list[T any] struct {
	buffer []T
}

// At equals to indexing on buffer slice.
func (l *list[T]) At(i int) T { return l.buffer[i] }

// Set equals to indexed assignment on buffer slice.
func (l *list[T]) Set(i int, item T) { l.buffer[i] = item }

// Len returns length of elements.
func (l *list[T]) Len() int { return len(l.buffer) }

// Cap returns capacity of buffer.
func (l *list[T]) Cap() int { return cap(l.buffer) }

// Empty reports length is zero.
func (l *list[T]) Empty() bool { return l.Len() == 0 }

func (l list[T]) String() string { return fmt.Sprint(l.buffer) }

// Clear removes all elements.
func (l *list[T]) Clear() { l.buffer = nil }

// Appends items to end of buffer.
func (l *list[T]) Append(items ...T) {
	l.buffer = append(l.buffer, items...)
}

// Clone returns immutable clone of list.
// Not copies capacity.
func (l *list[T]) Clone() List[T] {
	nl := new_list[T]()
	nl.buffer = make([]T, l.Len())
	_ = copy(nl.buffer, l.buffer)
	return nl
}

func new_list[T any]() *list[T] {
	l := new(list[T])
	l.buffer = nil
	return l;
}

// New returns new list instance.
func New[T any]() List[T] { return new_list[T]() }

// NewBuff returns new list instance by capacity.
//
// Special casees are;
//   NewBuff[T](cap) = cap accepts as zero (0) if cap < 0
func NewBuff[T any](cap int) List[T] {
	l := new_list[T]()
	if cap < 0 {
		cap = 0
	}
	l.buffer = make([]T, 0, cap)
	return l;
}
