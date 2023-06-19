package datastructures

import (
	"testing"
)

func TestInsertAndTop(t *testing.T) {
	var fh FibHeap[int]
	fh.Insert(10, 9, 2, 5, 2, -10)
	if fh.size != 6 {
		t.Error("Size check FAILED: got", fh.size)
	}
	if fh.Top() != -10 {
		t.Error("Top FAILED: Found", fh.Top())
	}
}

func TestPop(t *testing.T) {
	var fh FibHeap[int]
	fh.Insert(10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	t.Log("Top is ", fh.Top())
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range exp {
		if fh.Top() != v {
			t.Error("Pop FAILED: Top is ", fh.Top())
		}
		fh.Pop()
	}
}

func TestEmptyAndClearFH(t *testing.T) {
	var fh FibHeap[int]
	fh.Insert(10, 9, 2, 5, 2, -10)
	if fh.Empty() {
		t.Error("Empty FAILED, should be filled")
	}
	fh.Clear()
	if !fh.Empty() {
		t.Error("Empty FAILED, should be empty")
	}
}

func TestUnion(t *testing.T) {
	var fh1, fh2 FibHeap[int]
	fh1.Insert(10, 9, 2, 5, 2, 4)
	fh2.Insert(9, 2, 3, 5, 6, 1)
	fh1.Union(&fh2)
	exp := []int{1, 2, 2, 2, 3, 4, 5, 5, 6, 9, 9, 10}
	exp_size := 12
	for _, v := range exp {
		if fh1.size != exp_size || fh1.Top() != v {
			t.Error("Union FAILED: size is ", fh1.size, ", Top is ", fh1.Top())
		}
		fh1.Pop()
		exp_size--
	}
}
