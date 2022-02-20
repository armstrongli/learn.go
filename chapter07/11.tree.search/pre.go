package main

// todo 也可以使用for循环实现遍历
func preSearch(root *Node, targetNum int) bool {
	if root == nil {
		return false
	}
	totalCompare++
	if root.data == targetNum {
		return true
	}
	if preSearch(root.left, targetNum) {
		return true
	}
	if preSearch(root.right, targetNum) {
		return true
	}
	return false
}
