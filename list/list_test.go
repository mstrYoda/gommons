package list

import "testing"

func TestLen(t *testing.T) {
	l := New[int]()
	if l.Len() != 0 {
		t.Errorf("Len is should be zero")
	}

	const CAP = 20
	l2 := NewBuff[int](CAP)
	if l2.Len() != 0 {
		t.Errorf("Len is should be zero")
	}
}

func TestCap(t *testing.T) {
	l := New[int]()
	if l.Cap() != 0 {
		t.Errorf("Cap is should be zero")
	}

	const CAP = 20
	l2 := NewBuff[int](CAP)
	if l2.Cap() != CAP {
		t.Errorf("Cap is should be %d", CAP)
	}
}

func TestAppend(t *testing.T) {
	l := New[int]()
	l.Append(10, 20, 30, 40, 50)

	if l.Len() != 5 {
		t.Errorf("Len is should be 5")
	}
}
