package stream

import "sync"

// A Ring can be used as fixed size ring.
type Ring[T any] struct {
	elements []T
	index    int
	lock     sync.Mutex
}

// NewRing returns a Ring object with the given size n.
func NewRing[T any](n int) *Ring[T] {
	if n < 1 {
		panic("n should be greater than 0")
	}

	return &Ring[T]{
		elements: make([]T, n),
	}
}

// Add adds v into r.
func (r *Ring[T]) Add(v T) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.elements[r.index%len(r.elements)] = v
	r.index++
}

// Take takes all items from r.
func (r *Ring[T]) Take() []T {
	r.lock.Lock()
	defer r.lock.Unlock()

	var size int
	var start int
	if r.index > len(r.elements) {
		size = len(r.elements)
		start = r.index % len(r.elements)
	} else {
		size = r.index
	}

	elements := make([]T, size)
	for i := 0; i < size; i++ {
		elements[i] = r.elements[(start+i)%len(r.elements)]
	}

	return elements
}
