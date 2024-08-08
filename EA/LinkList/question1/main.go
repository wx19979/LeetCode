package main

import "fmt"

//链表的结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

//首先先定义一个构建链表的函数
func createList(list []int) *ListNode {
	//根据传入的数组list来构建链表
	head := &ListNode{Next: nil} //首先构建头指针
	//遍历list来进行插入操作
	for i := 0; i < len(list); i++ { //利用头插法初始化链表
		newNode := &ListNode{Val: list[i], Next: head.Next} //首先构造一个新的节点它的指针指向head的下一个节点
		head.Next = newNode                                 //然后头部指针的下一个节点指向它
	}
	return head //最终返回头部节点
}

//用于遍历的函数
func travelList(L *ListNode) {
	for L.Next != nil {
		fmt.Println(L.Next.Val) //输出结果
		L = L.Next              //继续向下遍历
	}
}

//创建用于更改指针的函数
func travelByN(L *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		L = L.Next
	}
	return L
}

// 删除链表中的节点函数
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val   //直接修改当前的值
	node.Next = node.Next.Next //然后将当前的下一个节点的值赋值到这里
}
func main() {
	list := []int{4, 5, 1, 9}
	ln := createList(list)
	travelList(ln)
	l := travelByN(ln, 3)
	deleteNode(l)
	fmt.Println()
	travelList(ln)
}
