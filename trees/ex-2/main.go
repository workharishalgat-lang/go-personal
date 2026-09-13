package main

func main() {
	// strict binary tree, small elements on strict left
	// large numbers on strict right

}
func BST(t *Tree, a int) {

	if a < t.data {
		if t.left == nil {
			t.left = &Tree{data: a}
		} else {
			BST(t.left, a)
		}
	} else {
		if t.right == nil {
			t.right = &Tree{data: a}
		} else {
			BST(t.right, a)
		}
	}

}

type Tree struct {
	data  int
	left  *Tree
	right *Tree
}
