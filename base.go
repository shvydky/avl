package avl

import (
	"iter"

	"golang.org/x/exp/constraints"
)

type Base[T constraints.Ordered] struct {
	root *node[T]
}

type node[T constraints.Ordered] struct {
	data   T
	left   *node[T]
	right  *node[T]
	height int
}

func New[T constraints.Ordered]() *Base[T] {
	return &Base[T]{
		root: nil,
	}
}

// Prev returns the previous element in the tree relative to n.
func (t *Base[T]) Prev(n T) (T, bool) {
	cur := t.root
	var prev *node[T]
	for cur != nil {
		if n < cur.data {
			cur = cur.left // Move to the left subtree
		} else if n > cur.data {
			prev = cur      // Update previous node
			cur = cur.right // Move to the right subtree
		} else {
			if cur.left != nil {
				return subtreeMax(cur.left).data, true // Return the maximum of the left subtree
			}
			if prev != nil {
				return prev.data, true // Return the previous node if it exists
			}
			return n, false // No previous element found
		}
	}
	if prev != nil {
		return prev.data, true // Return the previous node if it exists
	}
	return n, false // No previous element found
}

// Next returns the next element in the tree.
func (t *Base[T]) Next(n T) (T, bool) {
	cur := t.root
	var next *node[T]
	for cur != nil {
		if n < cur.data {
			next = cur     // Update next node
			cur = cur.left // Move to the left subtree
		} else if n > cur.data {
			cur = cur.right // Move to the right subtree
		} else {
			if cur.right != nil {
				return subtreeMin(cur.right).data, true // Return the minimum of the right subtree
			}
			if next != nil {
				return next.data, true // Return the next node if it exists
			}
			return n, false // No next element found
		}
	}
	if next != nil {
		return next.data, true // Return the next node if it exists
	}
	return n, false // No next element found
}

func (t *Base[T]) Traversal(from, to T) iter.Seq[T] {
	cur := newCursor(t.root, from, to)
	return func(yield func(T) bool) {
		for cur.Next() {
			v, ok := cur.Current()
			if !ok {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
	// if from < to {
	// 	return t.inOrderTraversal(t.root, from, to)
	// } else if from > to {
	// 	return t.reverseInOrderTraversal(t.root, from, to)
	// }
	// return func(yield func(T) bool) {}
}

// func (t *Base[T]) inOrderTraversal(n *node[T], from, to T) iter.Seq[T] {
// 	stack := make([]*node[T], 0, n.height)
// 	return func(yield func(T) bool) {
// 		current := n
// 		for current != nil || len(stack) > 0 {
// 			for current != nil {
// 				stack = append(stack, current)
// 				current = current.left
// 			}
// 			if len(stack) == 0 {
// 				break
// 			}
// 			current = stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]

// 			if current.data >= from && current.data <= to {
// 				if !yield(current.data) {
// 					return
// 				}
// 			}
// 			current = current.right
// 		}
// 	}
// }

// func (t *Base[T]) reverseInOrderTraversal(n *node[T], from, to T) iter.Seq[T] {
// 	stack := make([]*node[T], 0, n.height)
// 	return func(yield func(T) bool) {
// 		current := n
// 		for current != nil || len(stack) > 0 {
// 			for current != nil {
// 				stack = append(stack, current)
// 				current = current.right
// 			}
// 			if len(stack) == 0 {
// 				break
// 			}
// 			current = stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]

// 			if current.data <= from && current.data >= to {
// 				if !yield(current.data) {
// 					return
// 				}
// 			}
// 			current = current.left
// 		}
// 	}
// }

func (t *Base[T]) Insert(data T) (inserted bool) {
	node := &node[T]{
		data: data,
	}

	t.root, inserted = insert(t.root, node)
	return inserted
}

func (t *Base[T]) Remove(data T) (deleted bool) {
	t.root, deleted = remove(t.root, data)
	return deleted
}
