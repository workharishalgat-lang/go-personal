package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func addNode(node *Node) {
	// add content to the left
	// do you want to add node to the left
	fmt.Printf("do you want to add node to the left of %d press Y\n", node.data)

	var left string
	fmt.Scan(&left)
	if left == "Y" {
		fmt.Println("add the value which you want to add")
		var a int
		fmt.Scan(&a)
		node.left = &Node{data: a}
		addNode(node.left)
	}
	fmt.Printf("do you want to add node to the right of %d press Y\n", node.data)
	var right string
	fmt.Scan(&right)
	if right == "Y" {
		fmt.Println("add the value which you want to add")
		var a int
		fmt.Scan(&a)
		node.right = &Node{data: a}
		addNode(node.right)
	}
}

func main() {
	node := Node{data: 1}
	addNode(&node)
}
