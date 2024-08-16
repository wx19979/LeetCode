package main

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据数组创建二叉树(左子树是2i+1,右子树是2i+2)按照层序遍历创建的
func createBiTree(list []string) *TreeNode {
	//首先需要创建一个用于存储节点和索引的一个哈希表
	myTreeMap := map[int]*TreeNode{}
	if list[0] == "#" {
		return nil
	}
	//开始遍历操作加入索引表中
	for i := 0; i < len(list); i++ {
		if list[i] == "#" { //如果当前是空的直接赋值为空
			myTreeMap[i] = nil
		} else { //否则说明有数据
			tdata, _ := strconv.Atoi(list[i])                           //先转换一下
			myTreeMap[i] = &TreeNode{Left: nil, Right: nil, Val: tdata} //再开始进行赋值
		}
	}
	//从第二个节点开始都是有双亲节点的
	for i := 1; i < len(list); i++ {
		fm := (i - 1) / 2 //首先先找到双亲节点的索引
		if fm*2+1 == i {  //判断当前节点是双亲的左节点还是右节点
			myTreeMap[fm].Left = myTreeMap[i] //是左侧的就放到双亲节点的左面
		}
		if fm*2+2 == i { //如果是右侧节点
			myTreeMap[fm].Right = myTreeMap[i] //放置到双亲节点右侧
		}
	}
	return myTreeMap[0] //最终返回头一个节点的指针就是根节点
}

// 进行遍历操作(按照先根序遍历)
func Travel(bitTree *TreeNode) {
	if bitTree != nil { //如果当前节点不空
		fmt.Println(bitTree.Val) //输出结果
	}
	if bitTree.Left != nil {
		Travel(bitTree.Left) //递归遍历左子树
	}
	if bitTree.Right != nil {
		Travel(bitTree.Right) //递归遍历右子树
	}
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
	mytest1 := []string{"2", "1", "3"}
	testTree1 := createBiTree(mytest1)
	Travel(testTree1)
	isV1 := isValidBST1(testTree1)
	fmt.Println(isV1)
	mytest2 := []string{"5", "1", "4", "null", "null", "3", "6"}
	testTree2 := createBiTree(mytest2)
	Travel(testTree2)
	isV2 := isValidBST2(testTree2)
	fmt.Println(isV2)
}
