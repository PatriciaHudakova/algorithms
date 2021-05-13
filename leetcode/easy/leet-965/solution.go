/*
https://leetcode.com/problems/univalued-binary-tree/

Input: [1,1,1,1,1,null,1]
Output: true
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	treenode := createTreeNode()

	result := isUnivalTree(treenode)

	fmt.Println(result)
}

func isUnivalTree(root *TreeNode) bool {
	var leftCorrect bool
	var rightCorrect bool

	if root == nil {
		return true
	}

	if root.Left == nil || root.Val == root.Left.Val && isUnivalTree(root.Left) {
		leftCorrect = true
	} else {
		leftCorrect = false
	}

	if root.Right == nil || root.Val == root.Right.Val && isUnivalTree(root.Right) {
		rightCorrect = true
	} else {
		rightCorrect = false
	}

	return leftCorrect && rightCorrect
}

func createTreeNode() *TreeNode {
	treenode := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{1, nil, nil},
			Right: &TreeNode{1, nil, nil},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: &TreeNode{1, nil, nil},
		},
	}

	return &treenode
}
