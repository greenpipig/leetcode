package main

func isSymmetric(root *TreeNode) bool {
	return symmetricTrun(root.Left, root.Right)
}

func symmetricTrun(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return symmetricTrun(left.Left, right.Right) && symmetricTrun(left.Right, right.Left)
}
