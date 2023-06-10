package datastructures

import (
	"testing"
)

func TestLength(t *testing.T) {
	l := new(List[int])
	l.Append(1, 2, 3, 4, 5)
	if l.Length() != 5 {
		t.Error("Length FAILED")
	}
}

func TestEmpty(t *testing.T) {
	l := new(List[int])
	if !l.Empty() {
		t.Error("Empty FAILED: Should be empty")
	}
	l.Append(1)

	if l.Empty() {
		t.Error("Empty FAILED: Should not be empty")
	}

	l.Popfront()
	if !l.Empty() {
		t.Error("Empty FAILED: Should be empty")
	}
}

func TestElementAccess(t *testing.T) {
	l := new(List[int])
	l.Append(1, 2, 3, 4, 5)
	if l.Front() != 1 {
		t.Error("Front FAILED")
	}
	if l.Back() != 5 {
		t.Error("Back FAILED")
	}
}

func TestAppendPrepend(t *testing.T) {
	l := new(List[int])
	l.Append(4, 5, 6)
	l.Prepend(1, 2, 3)
	if l.Length() != 6 || l.Front() != 1 || l.Back() != 6 {
		t.Error("Append/Prepend FAILED")
	}
}

func TestFill(t *testing.T) {
	l := new(List[int])
	l.Fill(10, 1)
	if l.Length() != 10 {
		t.Error("Fill FAILED: Wrong length")
	}
	p := l.head
	for p != nil {
		if p.val != 1 {
			t.Error("Fill FAILED: Wrong content")
		}
		p = p.next
	}
}

func TestJoin(t *testing.T) {
	var l1, l2, ref List[int]
	l1.Append(1, 2, 3, 4, 5)
	l2.Append(6, 7, 8, 9, 10)
	ref.Append(1, 2, 6, 7, 8, 9, 10, 3, 4, 5)
	l1.Join(l2, 2)
	if l1.Length() != ref.Length() {
		t.Error("Join FAILED: l1 Length did not increase")
	}
	for lcur, refcur := l1.head, ref.head; lcur != nil; {
		if lcur.val != refcur.val {
			t.Error("Join FAILED: l1 element does not match reference")
		}
		lcur, refcur = lcur.next, refcur.next
	}
}

func TestPops(t *testing.T) {
	l := new(List[int])
	l.Append(1, 2, 3, 4, 5)
	l.Popfront()
	l.Popback()
	if l.Length() != 3 || l.Front() != 2 || l.Back() != 4 {
		t.Error("Pop FAILED")
	}
}

func TestErase(t *testing.T) {
	l := new(List[int])
	l.Append(1, 2, 3, 4, 5)
	l.Erase(2)
	if l.Length() != 4 {
		t.Error("Erase FAILED: Length did not decrease")
	}
	cur := l.head
	for cur != nil {
		if cur.val == 3 {
			t.Error("Erase FAILED: Erased element still in list")
		}
		cur = cur.next
	}
}

func TestClear(t *testing.T) {
	l := new(List[int])
	l.Append(1, 2, 3, 4, 5)
	l.Clear()
	if l.head != nil || l.Length() != 0 {
		t.Error("Clear FAILED")
	}
}

func TestComparisons(t *testing.T) {
	var l, leq, lle, lle_long, lgr, lgr_short List[int]
	l.Append(1, 2, 3, 4, 5)
	leq.Append(1, 2, 3, 4, 5)
	lle.Append(1, 2, 3, 4, 6)
	lle_long.Append(1, 2, 3, 4, 5, 6)
	lgr.Append(1, 2, 3, 4, 4)
	lgr_short.Append(1, 2, 3, 4)

	type tc struct {
		li  List[int]
		fun func(List[int]) bool
		err string
	}

	tests := []tc{
		{
			leq, l.Equal, "Equal FAILED",
		},
		{
			lle, l.Less, "Less FAILED",
		},
		{
			lle_long, l.Less, "Less longer FAILED",
		},
		{
			lgr, l.Greater, "Greater FAILED",
		},
		{
			lgr_short, l.Greater, "Greater shorter FAILED",
		},
	}

	for _, te := range tests {
		l.Print()
		te.li.Print()
		if !te.fun(te.li) {
			t.Error(te.err)
		}
	}
}
