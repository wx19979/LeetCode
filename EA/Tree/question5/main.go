package main

import (
	"math/rand"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将有序数组转换为二叉搜索树(利用递归法实现)
func sortedArrayToBST1(nums []int) *TreeNode {
	return helper1(nums, 0, len(nums)-1)
}

// 递归函数用于创建平衡二叉树
func helper1(nums []int, left, right int) *TreeNode {
	//实现判断左右参数是否有符合创建条件
	if left > right {
		return nil
	}
	//计算当前待插入节点所对应的下标位置
	mid := (left + right) / 2
	//创建当前的节点
	root := &TreeNode{Val: nums[mid]}
	//递归创建当前节点的左子树和右子树
	root.Left = helper1(nums, left, mid-1)
	root.Right = helper1(nums, mid+1, right)
	return root //最终返回根节点
}

// 递归函数用于创建平衡二叉树(随机法)
func sortedArrayToBST2(nums []int) *TreeNode {
	rand.Seed(time.Now().UnixNano())
	return helper2(nums, 0, len(nums)-1)
}

// 递归函数用于创建平衡二叉树(随机法)
func helper2(nums []int, left, right int) *TreeNode {
	//判断左右是否符合要求
	if left > right {
		return nil
	}
	//计算当前要创建节点的下标(这里的rand是随机函数在0和1之前选一个数)
	mid := (left + right + rand.Intn(2)) / 2
	//创建当前待插入的节点
	root := &TreeNode{Val: nums[mid]}
	//递归创建当前节点的左右孩子节点
	root.Left = helper2(nums, left, mid-1)
	root.Right = helper2(nums, mid+1, right)
	return root //最终返回根节点
}
func main() {

}
