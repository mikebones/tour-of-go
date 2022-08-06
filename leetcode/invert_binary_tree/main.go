package main

func main() {

}

type TreeNode struct {
	val int
	Right *TreeNode
	Left *TreeNode
}

func invertTree(treeNode *TreeNode) *TreeNode {
	if treeNode == nil {
		return nil
	}
	tmp := treeNode.Right
	treeNode.Right = treeNode.Left
	treeNode.Left = tmp
	invertTree(treeNode.Right)
	invertTree(treeNode.Left)
	return treeNode
}
