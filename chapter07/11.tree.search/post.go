package main

// todo 也可以使用for循环实现遍历
func postSearch(root *Node, targetNum int) bool {
	if root == nil {
		return false
	}
	if root.data > targetNum {
		if postSearch(root.left, targetNum) {
			return true
		}
	}
	if root.data < targetNum {
		if postSearch(root.right, targetNum) {
			return true
		}
	}
	totalCompare++
	if root.data == targetNum {
		return true
	}
	return false
}
