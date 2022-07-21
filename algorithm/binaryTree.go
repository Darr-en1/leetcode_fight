package algorithm

/*
二叉树是最容易培养框架思维的，而且大部分算法技巧，本质上都是树的遍历问题。

前序遍历：根结点 ---> 左子树 ---> 右子树

中序遍历：左子树---> 根结点 ---> 右子树

后序遍历：左子树 ---> 右子树 ---> 根结点

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// traverse 二叉树遍历模版
func traverse(root *TreeNode) {
	// 前序遍历
	traverse(root.Left)
	// 中序遍历
	traverse(root.Right)
	// 后序遍历
}

//  前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}

	// 1. 拆分成子问题
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	// 2. 子问题结果合并
	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)

	return res
}

//  中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	res = append(res, left...)
	res = append(res, root.Val)
	res = append(res, right...)

	return res
}

//  后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}

	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)

	res = append(res, left...)
	res = append(res, right...)
	res = append(res, root.Val)

	return res
}
