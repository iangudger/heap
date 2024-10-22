// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package heap implements a generic heap using the standard library's container/heap.
package heap

import "container/heap"

// A Comparable type can be compared with a method to other values of the same type.
type Comparable[T any] interface {
	// Less reports whether the receiver value must sort before the argument value.
	Less(v T) bool
}

// A Heap is a min-heap implemented as a slice with generic methods.
type Heap[T Comparable[T]] []T

// Init establishes the heap invariants required by the other routines in this package.
func (h *Heap[T]) Init() {
	heap.Init(h)
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
func (h *Heap[T]) Fix(i int) {
	heap.Fix(h, i)
}

// Len implements container/heap.Interface.Len and sort.Interface.Len.
//
// Since the Heap is defined to be a slice, using the built-in len is equivalent.
func (h *Heap[T]) Len() int {
	if h == nil {
		return 0
	}
	return len(*h)
}

// Less implements container/heap.Interface.Less and sort.Interface.Less.
func (h *Heap[T]) Less(i int, j int) bool {
	return (*h)[i].Less((*h)[j])
}

// Swap implements container/heap.Interface.Swap.
func (h *Heap[T]) Swap(i int, j int) {
	if i == j {
		return
	}
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// PushElement adds an element to the heap.
func (h *Heap[T]) PushElement(e T) {
	*h = append(*h, e)
	heap.Fix(h, len(*h)-1)
}

// MustPopElement removes and returns the min element in the heap. It panics if no elements are in the heap.
func (h *Heap[T]) MustPopElement() T {
	e := (*h)[0]
	i := h.Len() - 1
	(*h)[0] = (*h)[i]
	var zero T
	(*h)[i] = zero
	*h = (*h)[:i]
	heap.Fix(h, 0)
	return e
}

// PopElement removes and returns the min element in the heap.
func (h *Heap[T]) PopElement() (T, bool) {
	if h.Len() == 0 {
		var zero T
		return zero, false
	}
	return h.MustPopElement(), true
}

// Push implements container/heap.Interface.Push.
//
// Prefer PushElement over Push.
func (h *Heap[T]) Push(v any) {
	h.PushElement(v.(T))
}

// Pop implements container/heap.Interface.Pop.
//
// Prefer PopElement over Pop.
func (h *Heap[T]) Pop() any {
	return h.MustPopElement()
}

// MustPeekElement returns the min element in the heap. It panics if no elements are in the heap.
func (h *Heap[T]) MustPeekElement() T {
	return (*h)[0]
}

// PeekElement returns the min element in the heap.
func (h *Heap[T]) PeekElement() (T, bool) {
	if h.Len() == 0 {
		var zero T
		return zero, false
	}
	return (*h)[0], true
}

// RemoveElement removes and returns the element at index i from the heap.
func (h *Heap[T]) RemoveElement(i int) T {
	h.Swap(i, h.Len()-1)
	return h.MustPopElement()
}
