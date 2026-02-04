package aocutils

// and finally, gSets, a generic finite set implementation.
// Note that I'm setting it to use a struct{} because it uses
// zero bytes in memory. If I use bool, it will use one bit.

var exists = struct{}{}

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{
		elements: make(map[T]struct{}),
	}
	return s
}

func (s *Set[T]) Add(value T) {
	s.elements[value] = exists
}

func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, c := s.elements[value]
	return c
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.elements))
	for key := range s.elements {
		slice = append(slice, key)
	}
	return slice
}

// merge two sets
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.elements {
		result.Add(key)
	}
	for key := range other.elements {
		result.Add(key)
	}
	return result
}

// only elements in both sets
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.elements {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}
