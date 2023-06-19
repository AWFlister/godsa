package datastructures

// Fibonacci heap is a data structure that
// can retrieve the min or max of a collection quickly

import (
	"math"

	"golang.org/x/exp/constraints"
)

type FHNode[T constraints.Ordered] struct {
	val    T
	left   *FHNode[T]
	right  *FHNode[T]
	parent *FHNode[T]
	child  *FHNode[T]
	degree int
	mark   bool
}

type FibHeap[T constraints.Ordered] struct {
	min  *FHNode[T]
	size int
}

func (this *FibHeap[T]) Insert(v ...T) {
	for _, e := range v {
		eNode := FHNode[T]{e, nil, nil, nil, nil, 0, false}
		if this.min == nil {
			this.min = &eNode
			eNode.left = &eNode
			eNode.right = &eNode
		} else {
			eNode.left = this.min
			eNode.right = this.min.right
			eNode.right.left = &eNode
			eNode.left.right = &eNode
			if eNode.val < this.min.val {
				this.min = &eNode
			}
		}
		this.size++
	}
}

func (this *FibHeap[T]) Top() T {
	return this.min.val
}

func (this *FibHeap[T]) Pop() {
	z := this.min
	if z != nil {
		for z.child != nil {
			x := z.child
			if x.right == x {
				z.child = nil
			} else {
				z.child = x.right
			}
			x.right.left = x.left
			x.left.right = x.right
			x.right = z.right
			x.left = z
			z.right = x
			x.right.left = x
		}
		if z == z.right {
			this.min = nil
		} else {
			this.min = z.right
			this.min.left = z.left
			this.min.left.right = this.min
			this.consolidate()
		}
		this.size--
	}
}

func (this *FibHeap[T]) consolidate() {
	a := make([]*FHNode[T], int(math.Ceil(math.Log2(float64(this.size))))+1)
	for i := range a {
		a[i] = nil
	}
	x := this.min
	for a[x.degree] != x {
		d := x.degree
		for a[d] != nil && d < len(a) {
			y := a[d]
			if x.val > y.val {
				x, y = y, x
			}
			this.link(y, x)
			a[d] = nil
			d++
		}
		a[d] = x
		x = x.right
	}
	this.min = nil
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			if this.min == nil {
				this.min = a[i]
			} else if a[i].val <= this.min.val {
				this.min = a[i]
			}
		}
	}
}

func (this *FibHeap[T]) link(y, x *FHNode[T]) {
	y.left.right = y.right
	y.right.left = y.left
	if x.child == nil {
		x.child = y
		y.right = y
		y.left = y
	} else {
		w := x.child
		y.right = w.right
		y.left = w
		w.right.left = y
		w.right = y
	}
	x.degree++
}

func (this *FibHeap[T]) Union(fh *FibHeap[T]) {
	if fh.size > 0 {
		this.min.right.left = fh.min.left
		fh.min.left.right = this.min.right
		this.min.right = fh.min
		fh.min.left = this.min
		if fh.min.val < this.min.val {
			this.min = fh.min
		}
		this.size += fh.size
	}
}

func (this *FibHeap[T]) Empty() bool {
	return this.size == 0
}

func (this *FibHeap[T]) Clear() {
	this.min = nil
	this.size = 0
}

func (this *FibHeap[T]) Size() int {
	return this.size
}
