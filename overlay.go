package avl

import (
	"golang.org/x/exp/constraints"
)

type Overlay[T constraints.Ordered] struct {
	removed map[T]struct{}
	root    *node[T]
	base    Tree[T]
}

func NewOverlay[T constraints.Ordered](base Tree[T]) *Overlay[T] {
	return &Overlay[T]{
		removed: make(map[T]struct{}),
		root:    nil,
		base:    base,
	}
}

func (o *Overlay[T]) Insert(n T) (inserted bool) {
	o.Remove(n)
	o.root, inserted = insert(o.root, &node[T]{data: n})
	return inserted
}

func (o *Overlay[T]) Remove(n T) (removed bool) {
	if _, exists := o.removed[n]; !exists {
		o.removed[n] = struct{}{}
	}
	o.root, removed = remove(o.root, n)
	return removed
}

// func (o *Overlay[T]) Traversal(from, to T) iter.Seq[T] {
// 	baseit := o.base.Traversal(from, to)
// 	baseit()

// }

// func (o *Overlay[T]) Prev(n T) (T, bool) {}

// func (o *Overlay[T]) Next(n T) (T, bool) {}
