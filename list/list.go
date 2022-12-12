package list

import "fmt"

// List[T] is an interface for safe statically typing.
type List[T any] interface {
	At(int) T
	Set(int, T)
	Len() int
	Cap() int
	Empty() bool
	Append(...T)
	String() string
	Clear()
	Buffer() []T
	Insert(int, ...T)
	RemoveAt(i int)
	Map(func(T) T)
	Clone() List[T]
}

// list[T] is an wrapper for the Go slices.
type list[T any] struct { buff []T }

// At equals to indexing on buffer slice.
func (l *list[T]) At(i int) T { return l.buff[i] }

// Set equals to indexed assignment on buffer slice.
func (l *list[T]) Set(i int, item T) { l.buff[i] = item }

// Len returns length of elements.
func (l *list[T]) Len() int { return len(l.buff) }

// Cap returns capacity of buffer.
func (l *list[T]) Cap() int { return cap(l.buff) }

// Empty reports length is zero.
func (l *list[T]) Empty() bool { return l.Len() == 0 }

// Appends items to end of buffer.
func (l *list[T]) Append(items ...T) { l.buff = append(l.buff, items...) }

func (l list[T]) String() string { return fmt.Sprint(l.buff) }

// Clear removes all elements.
func (l *list[T]) Clear() { l.buff = nil }

// Buffer returns buffer (mutable).
func (l *list[T]) Buffer() []T { return l.buff }

// Insert inserts items by offset.
//
// Special cases are;
//   Insert(i, items) = appends if l.Empty() || i <= 0 || i >= l.Len()
func (l *list[T]) Insert(i int, items ...T) {
	// Special cases
	if l.Empty() || i <= 0 || i >= l.Len() {
		l.Append(items...)
		return
	}

	items = append(items, l.buff[i:]...)
	l.buff = append(l.buff[:i], items...)
}

// RemoveAt removes element by offset.
// Panics if i is out of range offset.
func (l *list[T]) RemoveAt(i int) {
	if l.Empty() || i < 0 || i >= l.Len() {
		panic("index is out of range")
	}
	l.buff = append(l.buff[:i], l.buff[i+1:]...)
}

// Map iterates on buffer and assign returned T to current offset.
// Panics if handler == nil.
func (l *list[T]) Map(handler func(T) T) {
	if handler == nil {
		panic("handler is nil")
	}
	for i := range l.buff {
		item := &l.buff[i]
		*item = handler(*item)
	}
}

// Clone returns immutable clone of list.
// Not copies capacity.
func (l *list[T]) Clone() List[T] {
	nl := new_list[T]()
	nl.buff = make([]T, l.Len())
	_ = copy(nl.buff, l.buff)
	return nl
}

func new_list[T any]() *list[T] {
	l := new(list[T])
	l.buff = nil
	return l;
}

// New returns new list instance.
func New[T any]() List[T] { return new_list[T]() }

// Make returns new list instance by length && capacity.
// Equals to make([]T, len, cap)
//
// Special casees are;
//   Make[T](len, cap) = len accepts as zero (0) if len < 0
//   Make[T](len, cap) = cap accepts as zero (0) if cap < 0
func Make[T any](len, cap int) List[T] {
	l := new_list[T]()
	if len < 0 {
		len = 0
	}
	if cap < 0 {
		cap = 0
	}
	l.buff = make([]T, len, cap)
	return l;
}
