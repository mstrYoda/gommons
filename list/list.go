package list

// List is an interface for safe statically typing.
type List[T any] interface {
	Len() int
	Cap() int
}

type list[T any] struct {
	buffer []T
}

// Len returns length of elements.
func (l *list[T]) Len() int { return len(l.buffer) }

// Cap returns capacity of buffer.
func (l *list[T]) Cap() int { return cap(l.buffer) }

func new_list[T any]() *list[T] {
	l := new(list[T])
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
