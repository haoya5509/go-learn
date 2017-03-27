package main

import "fmt"

type tree struct {
	value int
	left, right *tree
}

func sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	values = appendValues(values[:0], root)

	values = values[:len(values)-1]

	for _, v := range values {
		fmt.Println(v)
	}
}

func add(root *tree, value int) *tree {
	if root == nil {
		root = new(tree)
		root.value = value
		return root
	}

	if root.value > value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}

	return root
}

func appendValues(values []int, root *tree) []int {
	if root != nil {
		values = appendValues(values, root.left)
		values = append(values, root.value)
		values = appendValues(values, root.right)
	}
	return values
}

func main() {
	values := []int{3, 1, 7, 4, 2}
	sort(values[:4])

	fmt.Println(len(values), cap(values))
}