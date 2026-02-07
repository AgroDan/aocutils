package aocutils

// This is a deque (or double-ended queue) implementation,
// first realized that I needed it from advent of code 2018, day 8.

type Deque[T any] struct {
	items []T
}

func NewDeque[T any]() Deque[T] {
	return Deque[T]{}
}

// Shove everything in the front or top of the queue
func (d *Deque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

// Append to the end, or back of the queue
func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

// Pop the front or top of the queue
func (d *Deque[T]) PopFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}

	front := d.items[0]
	d.items = d.items[1:]

	return front, true
}

// Pop the back or end of the queue
func (d *Deque[T]) PopBack() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}

	back := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]

	return back, true
}

// Peek at the front of the queue without removing it
func (d *Deque[T]) PeekFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}

	return d.items[0], true
}

// Peek at the back of the queue without removing it
func (d *Deque[T]) PeekBack() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}

	return d.items[len(d.items)-1], true
}

// Check if the queue is empty
func (d *Deque[T]) IsEmpty() bool {
	return len(d.items) == 0
}

// Get the size of the queue
func (d *Deque[T]) Size() int {
	return len(d.items)
}

// Clear the queue
func (d *Deque[T]) Clear() {
	d.items = []T{}
}
