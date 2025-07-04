package avl

import "golang.org/x/exp/constraints"

func insert[T constraints.Ordered](root, node *node[T]) (res *node[T], inserted bool) {
	if root == nil {
		return node, true
	}
	if node.data < root.data {
		root.left, inserted = insert(root.left, node)
	} else if node.data > root.data {
		root.right, inserted = insert(root.right, node)
	}
	if !inserted {
		return root, false // Node already exists in the tree
	}

	root.height = 1 + max(height(root.left), height(root.right))
	return balance(root), true
}

func remove[T constraints.Ordered](root *node[T], data T) (res *node[T], deleted bool) {
	if root == nil {
		return nil, false // Node not found
	}
	if data < root.data {
		root.left, deleted = remove(root.left, data)
	} else if data > root.data {
		root.right, deleted = remove(root.right, data)
	} else {
		// Node found
		deleted = true
		if root.left == nil {
			return root.right, true // Node has no left child
		} else if root.right == nil {
			return root.left, true // Node has no right child
		}
		// Node has two children, find the in-order successor (smallest in the right subtree)
		successor := minValueNode(root.right)
		root.data = successor.data                         // Replace with the successor's data
		root.right, _ = remove(root.right, successor.data) // Delete the successor
	}
	root.height = 1 + max(height(root.left), height(root.right))
	return balance(root), deleted
}

func balance[T constraints.Ordered](n *node[T]) *node[T] {
	bf := balanceFactor(n)

	// Left Left Case
	if bf > 1 {
		// LR
		if n.left != nil && balanceFactor(n.left) < 0 {
			n.left = leftRotate(n.left)
		}
		// LL
		n = rightRotate(n)
	} else if bf < -1 {
		// RL
		if n.right != nil && balanceFactor(n.right) > 0 {
			n.right = rightRotate(n.right)
		}
		// RR
		n = leftRotate(n)
	}

	return n
}

func minValueNode[T constraints.Ordered](n *node[T]) *node[T] {
	for n != nil && n.left != nil {
		n = n.left // Traverse to the leftmost node
	}
	return n // Return the leftmost node
}

func balanceFactor[T constraints.Ordered](n *node[T]) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func leftRotate[T constraints.Ordered](n *node[T]) *node[T] {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.height = 1 + max(height(n.left), height(n.right))
	newRoot.height = 1 + max(height(newRoot.left), height(newRoot.right))

	return newRoot
}

func rightRotate[T constraints.Ordered](n *node[T]) *node[T] {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.height = 1 + max(height(n.left), height(n.right))
	newRoot.height = 1 + max(height(newRoot.left), height(newRoot.right))

	return newRoot
}

func height[T constraints.Ordered](n *node[T]) int {
	if n == nil {
		return 0
	}
	return n.height
}

func subtreeMax[T constraints.Ordered](n *node[T]) *node[T] {
	current := n
	if current == nil {
		return nil // Tree is empty
	}
	for current.right != nil {
		current = current.right // Traverse to the rightmost node
	}
	return current // Return the maximum node found
}

func subtreeMin[T constraints.Ordered](n *node[T]) *node[T] {
	current := n
	if current == nil {
		return nil // Tree is empty
	}
	for current.left != nil {
		current = current.left // Traverse to the leftmost node
	}
	return current // Return the minimum node found
}
