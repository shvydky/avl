package avl

import (
	"golang.org/x/exp/constraints"
)

type Cursor[T constraints.Ordered] interface {
	Current() (T, bool)
	Next() bool
	Prev() bool
}

type cursor[T constraints.Ordered] struct {
	stack      []*node[T]
	root       *node[T]
	current    *node[T]
	next, prev func() bool
	from, to   T
}

func newCursor[T constraints.Ordered](root *node[T], from, to T) Cursor[T] {
	c := &cursor[T]{
		root:  root,
		stack: make([]*node[T], 0, root.height),
		from:  from,
		to:    to,
	}
	if from < to {
		c.next = c.nextFn
		c.prev = c.prevFn
	} else {
		c.next = c.prevFn
		c.prev = c.nextFn
		c.from, c.to = c.to, c.from
	}

	return c
}

func (c *cursor[T]) Current() (T, bool) {
	if c.current == nil {
		var zero T
		return zero, false
	}
	return c.current.data, true
}

func (c *cursor[T]) reset() {
	for c.prev() {
	}
}

func (c *cursor[T]) nextFn() bool {
	if c.current == nil {
		c.current = c.root
		c.reset()
		return c.current != nil
	}

	if c.current.right != nil && c.current.right.data < c.to {
		c.stack = append(c.stack, c.current)
		c.current = c.current.right
		for c.current.left != nil && c.current.left.data > c.from {
			c.stack = append(c.stack, c.current)
			c.current = c.current.left
		}
		return true
	}

	pv := c.current.data
	for len(c.stack) > 0 {
		c.current = c.stack[len(c.stack)-1]
		c.stack = c.stack[:len(c.stack)-1]
		if c.current.data > pv {
			return true
		}
	}

	return false
}

func (c *cursor[T]) prevFn() bool {
	if c.current == nil {
		c.current = c.root
		c.reset()
		return c.current != nil
	}

	if c.current.left != nil && c.current.left.data >= c.from {
		c.stack = append(c.stack, c.current)
		c.current = c.current.left
		for c.current.right != nil && c.current.right.data < c.to {
			c.stack = append(c.stack, c.current)
			c.current = c.current.right
		}
		return true
	}

	pv := c.current.data
	for len(c.stack) > 0 {
		c.current = c.stack[len(c.stack)-1]
		c.stack = c.stack[:len(c.stack)-1]
		if c.current.data < pv {
			return true
		}
	}

	return false
}

func (c *cursor[T]) Next() bool {
	return c.next()
}

func (c *cursor[T]) Prev() bool {
	return c.prev()
}
