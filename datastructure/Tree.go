package datastructure

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}

	InOrder(root.Left)
	fmt.Println(root.Val)
	InOrder(root.Right)
}

func InOrderIterative(root *TreeNode) {
	if root == nil {
		return
	}

	var stack []*TreeNode
	current := root

	for {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else if len(stack) != 0 {
			n := len(stack)
			current = stack[n-1]
			stack = stack[:n-1]

			fmt.Println(current.Val)

			current = current.Right
		} else {
			break
		}
	}
}

// 	root := &datastructure.TreeNode{10, nil, nil}
//	root.Left = &datastructure.TreeNode{20, nil, nil}
//	root.Right = &datastructure.TreeNode{30, nil, nil}
//	root.Left.Left = &datastructure.TreeNode{40, nil, nil}
//	root.Left.Right = &datastructure.TreeNode{50, nil, nil}
//	root.Right.Right = &datastructure.TreeNode{60, nil, nil}
//	root.Left.Left.Left = &datastructure.TreeNode{70, nil, nil}
//
//	datastructure.InOrder(root)
