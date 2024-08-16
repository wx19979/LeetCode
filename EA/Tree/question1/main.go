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

// 利用递归的深度优先遍历搜索实现找到树的最大深度
func maxDepth1(root *TreeNode) int {
	//判断在空树的情况直接返回0
	if root == nil {
		return 0
	}
	//递归进行深度优先遍历
	return max(maxDepth1(root.Left), maxDepth1(root.Right)) + 1 //注意这里面没递归一层是需要加1的
}

// 首先创建一个比较大小的函数
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 利用队列的广度优先搜索实现找到树的最大深度
func maxDepth2(root *TreeNode) int {
	//首先判断树是否是空树
	if root == nil {
		return 0
	}
	//创建用于存储广度优先遍历的队列
	queue := []*TreeNode{}
	ans := 0 //创建用于存储最终结果的变量
	//将根节点放入队列中
	queue = append(queue, root)
	for len(queue) > 0 { //开始进行队列循环,队列完全为空循环结束
		sz := len(queue) //用于记录当前入队的元素个数
		for sz > 0 {     //开始对队列进行操作
			//首先取出当前遍历的节点
			node := queue[0]
			//然后对队列进行出队操作
			queue = queue[1:]
			//对计数器进行更新
			sz--
			//判断当前队列是否有子节点
			if node.Left != nil { //当前节点有左子树
				queue = append(queue, node.Left) //将左子树放到队列中
			}
			if node.Right != nil { //当前节点有右子树
				queue = append(queue, node.Right) //将右子树放入队列中
			}
		}
		ans++ //每次完成一层的操作计数器都加1记录层数
	}
	return ans //最终返回层数的结果
}

func main() {
	mytest := []string{"3", "9", "20", "#", "#", "15", "7"}
	testTree := createBiTree(mytest)
	Travel(testTree)
	i := maxDepth1(testTree)
	fmt.Println(i)
	j := maxDepth2(testTree)
	fmt.Println(j)
}
