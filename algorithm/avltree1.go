package algorithm

import "fmt"

// avl树  自平衡的二叉查找树

type AVLTree *AVLTreeNode

type AVLTreeNode struct {
	key   int
	high  int
	left  *AVLTreeNode
	right *AVLTreeNode
}

func AvlInsert(avl AVLTree, key int) AVLTree {
	if avl == nil {
		avl = new(AVLTreeNode)
		if avl == nil {
			fmt.Println("avl tree create error!")
			return nil
		} else {
			avl.key = key
			avl.high = 0
			avl.left = nil
			avl.right = nil
		}
	} else if key < avl.key {
		avl.left = AvlInsert(avl.left, key)
		if highTree(avl.left)-highTree(avl.right) == 2 {
			if key < avl.left.key { //LL
				avl = leftLeftRotation(avl)
			} else { // LR
				avl = leftRighRotation(avl)
			}
		}
	} else if key > avl.key {
		avl.right = AvlInsert(avl.right, key)
		if (highTree(avl.right) - highTree(avl.left)) == 2 {
			if key < avl.right.key { // RL
				avl = rightRightRotation(avl)
			} else {
				fmt.Println("right right", key)
				avl = rightLeftRotation(avl)
			}
		}
	} else if key == avl.key {
		fmt.Println("the key", key, "has existed!")
	}
	avl.high = max(highTree(avl.left), highTree(avl.right)) + 1
	return avl
}
func highTree(p AVLTree) int {
	if p == nil {
		return -1
	} else {
		return p.high
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func leftLeftRotation(k AVLTree) AVLTree {
	var kl AVLTree
	kl = k.left
	k.left = kl.right
	kl.right = k
	k.high = max(highTree(k.left), highTree(k.right)) + 1
	kl.high = max(highTree(kl.left), k.high) + 1
	return kl
}

func rightRightRotation(k AVLTree) AVLTree {
	var kr AVLTree
	kr = k.right
	k.right = kr.left
	kr.left = k
	k.high = max(highTree(k.left), highTree(k.right)) + 1
	kr.high = max(k.high, highTree(kr.right)) + 1
	return kr
}

func leftRighRotation(k AVLTree) AVLTree {
	k.left = rightRightRotation(k.left)
	return leftLeftRotation(k)
}

func rightLeftRotation(k AVLTree) AVLTree {
	k.right = leftLeftRotation(k.right)
	return rightRightRotation(k)
}

func PreOrder(avl AVLTree) {
	if avl != nil {
		fmt.Print(avl.key, "\t")
		PreOrder(avl.left)
		PreOrder(avl.right)
	}

}

/**
前序非递归遍历
*/
func PreNotRecursionOrder(root AVLTree) []int {
	if root == nil {
		return nil
	}
	var stack *ArrayStack
	array := make([]int, 0, 0)
	for root != nil || !IsEmpty(stack) {
		if root != nil {
			array = append(array, root.key)
			stack = Push(stack, root)
			root = root.left
		} else {
			root = Pop(stack).(*AVLTreeNode)
			root = root.right
		}
	}
	return array
}

/**
中序遍历递归
*/
func MidOrder(avl AVLTree) {
	if avl != nil {
		MidOrder(avl.left)
		fmt.Print(avl.key, "\t")
		MidOrder(avl.right)
	}
}

/**
后续遍历递归
*/
func PoOrder(avl AVLTree) {
	if avl == nil {
		return
	}
	PoOrder(avl.left)
	PoOrder(avl.right)
	fmt.Println(avl.key)
}

func getAvlHeight(avl AVLTree) int {
	if avl == nil {
		return 0
	}
	l := getAvlHeight(avl.left) + 1
	r := getAvlHeight(avl.right) + 1
	if l > r {
		return l
	}
	return r
}
