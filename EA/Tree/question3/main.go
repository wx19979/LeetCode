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

// 用于查看是否是对称的树的函数,利用递归的方式进行检查
func isSymmetric1(root *TreeNode) bool {
	return check(root, root)
}

// 用于具体的检查操作
func check(p, q *TreeNode) bool {
	if p == nil && q == nil { //如果当前都是空的说明是对称的
		return true
	}
	if p == nil || q == nil { //只有一个有节点说明不对称
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left) //递归判断两个值是否是相等
}

// 用于查看是否是对称的树的函数,利用迭代法和队列进行求解
func isSymmetric2(root *TreeNode) bool {
	//创建用于存储对比左右子树的节点变量
	u, v := root, root
	//创建用于进行遍历操作的队列
	q := []*TreeNode{}
	q = append(q, u) //将左右都加入队列中
	q = append(q, v)
	for len(q) > 0 { //开始循环遍历
		u, v = q[0], q[1] //从队列中取出两元素
		q = q[2:]         //进行出队操作
		//进行判断操作
		if u == nil && v == nil { //如果两个节点都是空的就继续向下遍历进行对比
			continue
		}
		if u == nil || v == nil { //如果其中有一个是空的说明不对称直接返回错误
			return false
		}
		if u.Val != v.Val { //如果两个值不同说明也是不对称的
			return false
		}
		//将树的节点按照应该对称的方式加入到队列中以使得能够继续进行判断树是否对称
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true //整体比对都没有错误直接返回正确信息
}

func main() {
	mytest1 := []string{"1", "2", "2", "3", "4", "4", "3"}
	testTree1 := createBiTree(mytest1)
	Travel(testTree1)
	isS1 := isSymmetric1(testTree1)
	fmt.Println(isS1)
	mytest2 := []string{"1", "2", "2", "#", "3", "#", "3"}
	testTree2 := createBiTree(mytest2)
	Travel(testTree2)
	isS2 := isSymmetric2(testTree2)
	fmt.Println(isS2)
}
