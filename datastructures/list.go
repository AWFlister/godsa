package datastructures

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// List represents a doubly-linked list that holds
// values of any type.

// Type definitions
type ListNode[T constraints.Ordered] struct {
	prev *ListNode[T]
	next *ListNode[T]
	val  T
}

type List[T constraints.Ordered] struct {
	head *ListNode[T]
	tail *ListNode[T]
	size int
}

// Capacity
func (l *List[T]) Length() int {
	return l.size
}

func (l *List[T]) Empty() bool {
	return l.size == 0
}

// Element Access
func (l *List[T]) Front() T {
	return l.head.val
}

func (l *List[T]) Back() T {
	return l.tail.val
}

// Modifiers
func (l *List[T]) Append(val ...T) {
	for _, v := range val {
		node := ListNode[T]{nil, nil, v}
		if l.head == nil {
			l.head = &node
			l.tail = &node
		} else {
			l.tail.next = &node
			node.prev = l.tail
			l.tail = &node
		}
		l.size++
	}
}

func (l *List[T]) Prepend(val ...T) {
	for i := range val {
		node := ListNode[T]{nil, nil, val[len(val)-i-1]}
		if l.head == nil {
			l.head = &node
			l.tail = &node
		} else {
			node.next = l.head
			l.head.prev = &node
			l.head = &node
		}
		l.size++
	}
}

func (l *List[T]) Fill(len int, elem T) {
	for ; len > 0; len-- {
		l.Append(elem)
	}
}

func (l *List[T]) Join(src List[T], pos int) {
	if pos >= 0 && pos <= l.Length() && !src.Empty() {
		lcur := l.head
		for ; pos > 1; pos-- {
			lcur = lcur.next
		}
		if l.Empty() {
			l.head = src.head
			l.tail = src.tail
		} else if pos == 0 {
			src.tail.next = l.head
			src.tail.next.prev = src.tail
			l.head = src.head
		} else if lcur == l.tail {
			l.tail.next = src.head
			l.tail.next.prev = l.tail
			l.tail = src.tail
		} else {
			lcur.next.prev = src.tail
			src.tail.next = lcur.next
			lcur.next = src.head
			src.head.prev = lcur
		}
		l.size += src.Length()
	}
}

func (l *List[T]) Popback() {
	if l.size > 0 {
		l.tail = l.tail.prev
		l.size--
	}
}

func (l *List[T]) Popfront() {
	if l.size > 0 {
		l.head = l.head.next
		l.size--
	}
}

func (l *List[T]) Erase(pos int) {
	if pos >= 0 && pos < l.size {
		ptr := l.head
		for ; pos > 0; pos-- {
			ptr = ptr.next
		}
		if ptr.prev != nil {
			ptr.prev.next = ptr.next
		}
		if ptr.next != nil {
			ptr.next.prev = ptr.prev
		}
		l.size--
	}
}

func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Comparisons
func (l *List[T]) Equal(li List[T]) bool {
	h1, h2 := l.head, li.head
	for {
		if (h1 == nil) || (h2 == nil) {
			return h1 == h2
		}
		if h1.val != h2.val {
			return false
		}
		h1 = h1.next
		h2 = h2.next
	}
}

func (l *List[T]) Less(li List[T]) bool {
	h1, h2 := l.head, li.head
	for {
		if (h1 == nil) || (h2 == nil) {
			return h1 == nil
		}
		if h1.val < h2.val {
			return true
		} else if h1.val > h2.val {
			return false
		}
		h1 = h1.next
		h2 = h2.next
	}
}

func (l *List[T]) Greater(li List[T]) bool {
	h1, h2 := l.head, li.head
	for {
		if (h1 == nil) || (h2 == nil) {
			return h2 == nil
		}
		if h1.val > h2.val {
			return true
		} else if h1.val < h2.val {
			return false
		}
		h1 = h1.next
		h2 = h2.next
	}
}

func (l *List[T]) LessOrEqual(li List[T]) bool {
	h1, h2 := l.head, li.head
	for {
		if (h1 == nil) || (h2 == nil) {
			return h1 == h2
		}
		if h1.val < h2.val {
			return true
		} else if h1.val > h2.val {
			return false
		}
		h1 = h1.next
		h2 = h2.next
	}
}

func (l *List[T]) GreaterOrEqual(li List[T]) bool {
	h1, h2 := l.head, li.head
	for {
		if (h1 == nil) || (h2 == nil) {
			return h1 == h2
		}
		if h1.val > h2.val {
			return true
		} else if h1.val < h2.val {
			return false
		}
		h1 = h1.next
		h2 = h2.next
	}
}

func (l *List[T]) Print() {
	cur := l.head
	fmt.Print("[")
	for {
		fmt.Print(cur.val)
		if cur == l.tail {
			break
		} else {
			fmt.Print(" ")
			cur = cur.next
		}
	}
	fmt.Println("]")
}
