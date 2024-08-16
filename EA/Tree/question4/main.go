package main

import (
	"fmt"
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

// 对树进行层序遍历的函数
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

func main() {
	mytest := []string{"3", "9", "20", "#", "#", "15", "7"}
	testTree := createBiTree(mytest)
	Travel(testTree)
	res := levelOrder(testTree)
	fmt.Println(res)
}
