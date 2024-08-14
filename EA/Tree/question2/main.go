package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 利用递归法判断当前的树是否是二叉搜索树
func isValidBST1(root *TreeNode) bool {
	return helper(root, math.MaxInt64, math.MinInt64) //返回调用这个函数
}

// 创建一个辅助的函数
func helper(root *TreeNode, upper, lower int) bool {
	if root == nil { //如果当前是空的
		return true //直接返回正确信息
	}
	if root.Val <= lower || root.Val >= upper { //如果不满足条件
		return false //直接返回错误信息
	}
	return helper(root.Left, root.Val, lower) && helper(root.Right, upper, root.Val) //递归调用返回左右子树两者的结果
}
func isValidBST2(root *TreeNode) bool {
	//首先创建一个栈用于存储中序遍历的内容
	stack := []*TreeNode{}
	//创建一个用于存储比较数字大小的变量
	inorder := math.MinInt64 //先设置为int的最小值
	//开始进行中序遍历操作
	for len(stack) > 0 || root != nil {
		//根据中序遍历顺序开始先遍历左子树
		for root != nil {
			stack = append(stack, root) //将遍历到的节点内容放入栈中
			root = root.Left            //逐步向左遍历一直遍历到最左下角
		}
		//开始遍历根节点
		root = stack[len(stack)-1]   //找到栈顶元素,找到理论上该树的最小值
		stack = stack[:len(stack)-1] //出栈操作,将栈顶元素抛弃
		if root.Val <= inorder {     //将理论上最小值和比较的变量进行比较如果不满足条件
			return false //直接返回错误信息
		}
		inorder = root.Val //更新比较的变量
		root = root.Right  //开始遍历右子树
	}
	return true //整体遍历都没有错误直接返回正确的信息
}
func main() {

}
