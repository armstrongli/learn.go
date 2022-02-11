package main

import (
	"fmt"
)

type LinkNode struct {
	data int
	next *LinkNode
}

func main() {
	n1 := &LinkNode{
		data: 1,
		next: nil,
	}
	n2 := &LinkNode{
		data: 2,
		next: nil,
	}
	n3 := &LinkNode{
		data: 3,
		next: nil,
	}
	n4 := &LinkNode{
		data: 4,
		next: nil,
	}
	n6 := &LinkNode{
		data: 6,
		next: nil,
	}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6

	{
		rangeLink(n1)
	}

	{
		fmt.Println("插入5")
		n5 := &LinkNode{
			data: 5,
			next: nil,
		}
		insertNode(n1, n5)
		insertNode(n1, &LinkNode{
			data: 7,
			next: nil,
		})
		insertNode(n1, &LinkNode{
			data: 5,
			next: nil,
		})
		insertNode(n1, &LinkNode{
			data: 3,
			next: nil,
		})
		rangeLink(n1)
	}

	{
		fmt.Println("删除节点")
		n1 = deleteNode(n1, 3)
		n1 = deleteNode(n1, 5)
		n1 = deleteNode(n1, 1)
		rangeLink(n1)
	}
}

func deleteNode(root *LinkNode, data int) *LinkNode {
	tmpNode := root
	if root != nil && root.data == data {
		if root.next == nil {
			return nil
		}
		right := root.next
		tmpNode.next = nil
		return right
	}
	for {
		if tmpNode.next == nil {
			break
		}
		right := tmpNode.next
		if right.data == data {
			// 找到要删除的节点，开始删除
			tmpNode.next = right.next
			right.next = nil
			return root
		}
		tmpNode = tmpNode.next
	}
	return root
}

func insertNode(root *LinkNode, newNode *LinkNode) {
	tmpNode := root
	for {
		if tmpNode != nil {
			if newNode.data > tmpNode.data {
				if tmpNode.next == nil {
					// 已经到结尾，直接追加
					tmpNode.next = newNode
				} else {
					if tmpNode.next.data >= newNode.data {
						// 找到合适位置，准备插入数据
						newNode.next = tmpNode.next
						tmpNode.next = newNode
						break
					}
				}
			}
		} else {
			break
		}
		tmpNode = tmpNode.next
	}
}

func rangeLink(root *LinkNode) {
	tmpNode := root
	for {
		if tmpNode != nil {
			fmt.Println(tmpNode.data)
		} else {
			break
		}
		tmpNode = tmpNode.next
	}
}
