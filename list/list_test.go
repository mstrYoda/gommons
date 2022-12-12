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

func TestEmpty(t *testing.T) {
	l := New[int]()

	if !l.Empty() {
		t.Errorf("Empty should return true")
	}

	l.Append(10, 20, 30, 40, 50)

	if l.Empty() {
		t.Errorf("Empty should return false")
	}
}

func TestString(t *testing.T) {
	l := New[int]()

	if l.String() != "[]" {
		t.Errorf("String should return []")
	}

	l.Append(10, 20, 30, 40, 50)

	if l.String() != "[10 20 30 40 50]" {
		t.Errorf("Empty should return [10 20 30 40 50]")
	}
}

func TestClear(t *testing.T) {
	l := New[int]()
	l.Append(10, 20, 30, 40, 50)
	l.Clear()

	if l.Len() != 0 {
		t.Errorf("Len should return zero")
	}
}

func TestClone(t *testing.T) {
	l := New[int]()
	l.Append(10, 20, 30, 40, 50)
	
	l2 := l.Clone()

	if l.String() != l2.String() {
		t.Errorf("Clone returned false clone")
	}
}
