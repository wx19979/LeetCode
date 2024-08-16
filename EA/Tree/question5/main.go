package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	//首先创建用于存储最终结果的二维数组
	res := [][]int{}
	if root == nil { //判断是否是空树
		return res //空树直接返回空值
	}
	//创建一个以根节点为开头队列
	arr := []*TreeNode{root}
	//开始循环操作
	for len(arr) > 0 {
		//首先重置长度
		size := len(arr)
		//创建一个用于存储当前层的结果的数组
		curRes := []int{}
		//开始遍历队列
		for i := 0; i < size; i++ {
			//首先从队列取出当前遍历的元素
			node := arr[i]
			curRes = append(curRes, node.Val) //将当前遍历的值放到暂存结果的数组
			//如果左右子树都存在将左右子树也放到当前的队列中
			if node.Left != nil {
				arr = append(arr, node.Left)
			}
			if node.Right != nil {
				arr = append(arr, node.Right)
			}
		}
		//将刚刚都遍历过的元素进行出队处理
		arr = arr[size:]
		//该层便利完将获得的暂存结果放到结果二维数组中
		res = append(res, curRes)
	}
	//最终返回结果
	return res
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
	nums1 := []int{-10, -3, 0, 5, 9}
	tree1 := sortedArrayToBST2(nums1)
	res1 := levelOrder(tree1)
	fmt.Println(res1)
	nums2 := []int{1, 3}
	tree2 := sortedArrayToBST1(nums2)
	res2 := levelOrder(tree2)
	fmt.Println(res2)
}
