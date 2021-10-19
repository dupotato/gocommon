package algorithm

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
获取二叉树的最小深度
*/
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(*root)
	minDepth := 0
	for queue.Len() > 0 {
		newQueue := list.New()
		minDepth++
		for queue.Len() > 0 {
			e := queue.Front()
			queue.Remove(e)
			tempTree := e.Value.(TreeNode)
			if tempTree.Left == nil && tempTree.Right == nil {
				return minDepth
			}
			if tempTree.Left != nil {
				newQueue.PushBack(*tempTree.Left)
			}
			if tempTree.Right != nil {
				newQueue.PushBack(*tempTree.Right)
			}
		}
		queue.PushBackList(newQueue)
	}
	return minDepth
}

func LevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0, 0)
	res = order(root, res, 0)
	return res
}

func order(root *TreeNode, res [][]int, level int) [][]int {

	if root == nil {
		return res
	}
	if len(res)-1 < level {
		res = append(res, make([]int, 0, 0))
	}
	res[level] = append(res[level], root.Val)
	res = order(root.Left, res, level+1)
	res = order(root.Right, res, level+1)
	return res
}

//preorderTransvel
var res []int
var stack []*TreeNode

func preorderTransvel(root *TreeNode) []int {

	for root != nil || len(stack) > 0 {

		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) > 0 {
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}

	return res
}

func inorderTransver(root *TreeNode) {
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
	}
}
