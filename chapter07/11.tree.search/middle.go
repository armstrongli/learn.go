package main

// todo 也可以使用for循环实现遍历
func midSearch(root *Node, targetNum int) bool {
	if root == nil {
		return false
	}
	if midSearch(root.left, targetNum) {
		return true
	}
	totalCompare++
	if root.data == targetNum {
		return true
	}
	if midSearch(root.right, targetNum) {
		return true
	}
	return false
}
