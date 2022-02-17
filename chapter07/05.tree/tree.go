package main

import (
	"fmt"
)

type Node struct {
	data  int
	root  *Node
	left  *Node
	right *Node
}

func buildTree() *Node {
	n1 := &Node{data: 51}
	n2 := &Node{data: 35}
	n3 := &Node{data: 65}

	n1.left = n2
	n2.root = n1
	n1.right = n3
	n3.root = n1
	return n1
}

func insertNode(root *Node, newNode *Node) *Node {
	if root == nil {
		return newNode
	}
	if newNode.data == root.data {
		return root
	}
	if newNode.data < root.data {
		if root.left == nil {
			// 找到位置，插入数据
			root.left = newNode
			newNode.root = root
		} else {
			insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			// 找到位置，插入数据
			root.right = newNode
			newNode.root = root
		} else {
			insertNode(root.right, newNode)
		}
	}
	return root
}

// 中间用于帮助理解删除的是最后叶子节点的情况
func deleteNodeLeaf(root *Node, v int) *Node {
	leftRoot := root
	if leftRoot.data == v && leftRoot.left == nil && leftRoot.right == nil {
		leftRoot = leftRoot.root
		right := root
		if leftRoot.left == right {
			// 删除左边叶子
			leftRoot.left = nil
			right.root = nil
			return leftRoot
		} else {
			// 删除右边叶子
			leftRoot.right = nil
			right.root = nil
			return leftRoot
		}
	}
	return root
}

func deleteNode(root *Node, v int) *Node {
	if v < root.data {
		deleteNode(root.left, v)
	} else if v > root.data {
		deleteNode(root.right, v)
	} else {
		// 现在root指向要删除的节点
		leftNextGen := findNextGenFromLeft(root.left)
		rightNextGen := findNextGenFromRight(root.right)
		if leftNextGen == nil && rightNextGen == nil {
			// 现在要删除的是叶子结点，即最底部的节点
			top := root.root
			down := root
			if top.left == down {
				// 表示是左子树
				top.left = nil
				down.root = nil
				return nil
			} else {
				// 表示是右子树
				top.right = nil
				down.root = nil
				return nil
			}
		} else if leftNextGen != nil {
			root.data = leftNextGen.data
			deleteNode(leftNextGen, leftNextGen.data)
		} else if rightNextGen != nil {
			root.data = rightNextGen.data
			deleteNode(rightNextGen, rightNextGen.data)
		}
	}
	return root
}

func findNextGenFromLeft(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.right != nil {
			tmpNode = tmpNode.right
		} else {
			break
		}
	}
	return tmpNode
}

func findNextGenFromRight(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.left != nil {
			tmpNode = tmpNode.left
		} else {
			break
		}
	}
	return tmpNode
}

func main() {
	root := buildTree()
	insertNode(root, &Node{data: 43})
	insertNode(root, &Node{data: 28})
	fmt.Println("done")

	fmt.Println("删除Node")
	deleteNode(root, 43)
	deleteNode(root, 35)
	fmt.Println("删除结束")
}
