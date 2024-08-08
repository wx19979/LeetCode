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

//删掉倒数第n个元素(利用的是快慢指针)
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	//首先创建一个哑节点,它处在头节点前面
	dummy := &ListNode{0, head}
	//创建快慢指针分别指向头节点和哑节点
	fast, slow := head, dummy
	//先将快指针向后移动n次
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	//然后开始继续移动快指针
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//最后修改slow的指针,让其指向下一个位置
	slow.Next = slow.Next.Next
	return dummy.Next //最后返回哑节点的下一个节点,也就是头节点
}

//用于获取链表长度的函数
func getListLen(head *ListNode) int {
	p, i := head, 0
	for p.Next != nil {
		i++
		p = p.Next
	}
	return i
}

//删掉倒数第n个元素(利用的是计数方法)
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	p, j := dummy, getListLen(head)-n //设置一个用于遍历的指针,并且获取到前面需要移动的次数
	for i := 0; i <= j; i++ {
		p = p.Next //p根据次数不断向后移动
	}
	p.Next = p.Next.Next
	return dummy.Next
}

//删掉倒数第n个元素(利用的是栈)
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	//首先创建一个用于存储内容的栈
	nodes := []*ListNode{}
	//创建一个哑节点
	dummy := &ListNode{0, head}
	//创建用于操作的指针
	p := dummy
	//开始循环操作将链表的内容放入栈中
	for p != nil {
		nodes = append(nodes, p)
		p = p.Next
	}
	//然后根据栈的特性找到对应位置的节点
	prev := nodes[len(nodes)-n-1]
	prev.Next = prev.Next.Next //修改指针
	return dummy.Next          //返回头节点
}
func main() {
	list1 := []int{1, 2}
	ln1 := createList(list1)
	travelList(ln1)
	fmt.Println()
	ln1 = removeNthFromEnd1(ln1, 1)
	travelList(ln1)
	fmt.Println()
	list2 := []int{1, 2}
	ln2 := createList(list2)
	travelList(ln2)
	fmt.Println()
	ln2 = removeNthFromEnd2(ln2, 1)
	travelList(ln2)
	fmt.Println()
	list3 := []int{1, 2}
	ln3 := createList(list3)
	travelList(ln3)
	fmt.Println()
	ln3 = removeNthFromEnd3(ln3, 1)
	travelList(ln3)
	fmt.Println()
}
