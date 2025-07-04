package avl

import (
	"iter"

	"golang.org/x/exp/constraints"
)

type Tree[T constraints.Ordered] interface {
	// Insert adds a new element to the tree, returning true if the element was added.
	Insert(value T) bool
	// Delete removes an element from the tree, returning true if the element was removed.
	Delete(value T) bool
	// Traversal returns an iterator for elements in the range [from, to].
	Traversal(from, to T) iter.Seq[T]
	// Cursor creates a cursor for the tree, allowing traversal from 'from' to 'to'.
	Cursor(from, to T) Cursor[T]
}
