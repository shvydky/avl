package avl

import (
	"slices"
	"testing"
)

func Test_Insert(t *testing.T) {
	tree := New[int]()
	if !tree.Insert(10) {
		t.Error("expected to insert 10")
	}
	if !tree.Insert(20) {
		t.Error("expected to insert 20")
	}
	if !tree.Insert(5) {
		t.Error("expected to insert 5")
	}
	if tree.root.data != 10 || tree.root.left.data != 5 || tree.root.right.data != 20 {
		t.Error("tree structure is incorrect after inserts")
	}
	if tree.root.height != 1 {
		t.Error("expected height of root to be 2")
	}
}

func Test_Delete(t *testing.T) {
	tree := New[int]()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(30)
	tree.Insert(40)

	if !tree.Remove(20) {
		t.Error("expected to delete 20")
	}

	if !tree.Remove(10) {
		t.Error("expected to delete 10")
	}

	if !tree.Remove(5) {
		t.Error("expected to delete 5")
	}

	if !tree.Remove(30) {
		t.Error("expected to delete 30")
	}

	if !tree.Remove(40) {
		t.Error("expected to delete 40")
	}

	if tree.root != nil {
		t.Error("expected tree to be empty after deleting all nodes")
	}
}

func Test_Traversal(t *testing.T) {
	tree := New[int]()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(0)
	tree.Insert(-5)
	tree.Insert(50)
	tree.Insert(40)

	expected := []int{5, 10, 20}

	i := 0
	for v := range tree.Traversal(3, 25) {
		if i >= len(expected) {
			t.Errorf("unexpected value %d at index %d", v, i)
			return
		}
		if v != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, v)
		}
		i++
	}
	if i != len(expected) {
		t.Errorf("expected %d values, got %d", len(expected), i)
	}

	slices.Reverse(expected)

	// i = 0
	// for v := range tree.Traversal(25, 3) {
	// 	if i >= len(expected) {
	// 		t.Errorf("unexpected value %d at index %d", v, i)
	// 		return
	// 	}
	// 	if v != expected[i] {
	// 		t.Errorf("expected %d at index %d, got %d", expected[i], i, v)
	// 	}
	// 	i++
	// }
	// if i != len(expected) {
	// 	t.Errorf("expected %d values, got %d", len(expected), i)
	// }
}

func Test_PrevNext(t *testing.T) {
	tree := New[int]()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(0)
	tree.Insert(-5)
	tree.Insert(50)
	tree.Insert(40)

	prev, ok := tree.Prev(10)
	if !ok || prev != 5 {
		t.Errorf("expected previous of 10 to be 5, got %d", prev)
	}

	next, ok := tree.Next(10)
	if !ok || next != 20 {
		t.Errorf("expected next of 10 to be 20, got %d", next)
	}

	prev, ok = tree.Prev(5)
	if !ok || prev != 0 {
		t.Errorf("expected previous of 5 to be 0, got %d", prev)
	}

	_, ok = tree.Next(50)
	if ok {
		t.Error("expected no next for 50")
	}
}
