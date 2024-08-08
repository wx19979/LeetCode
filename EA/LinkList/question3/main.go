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
	head := &ListNode{list[len(list)-1], nil} //首先构建头指针
	//遍历list来进行插入操作
	for i := 0; i < len(list)-1; i++ { //利用头插法初始化链表
		newNode := &ListNode{Val: list[i], Next: head.Next} //首先构造一个新的节点它的指针指向head的下一个节点
		head.Next = newNode                                 //然后头部指针的下一个节点指向它
	}
	return head //最终返回头部节点
}

//用于遍历的函数
func travelList(L *ListNode) {
	for L != nil {
		fmt.Println(L.Val) //输出结果
		L = L.Next         //继续向下遍历
	}
}

//翻转链表函数(利用栈进行操作),这道题有坑head里面是带有数值的,所以想做到和答案相同需要加一个哑巴节点
func reverseList1(head *ListNode) *ListNode {
	nodes := []*ListNode{}      //创建一个存储节点的数组
	dummy := &ListNode{0, head} //创建一个哑巴节点
	p := dummy.Next             //获取到头节点,从头节点开始向后依次遍历
	for p != nil {              //开始逐个遍历并且放入栈中
		nodes = append(nodes, p)
		p = p.Next
	}
	dummy.Next = nil                  //开始重建指向关系
	for i := 0; i < len(nodes); i++ { //循环进行翻转
		nodes[i].Next = dummy.Next //修改指针的指向
		dummy.Next = nodes[i]
	}
	return dummy.Next //最终返回哑巴节点的后面的节点即可
}

//翻转链表函数(利用迭代)
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode //首先创建一个工具指针指向前一个节点
	curr := head       //再创建一个指向当前的节点的操作指针
	//开始循环改变指针方向的操作
	for curr != nil {
		next := curr.Next //首先获取到当前节点的下一个节点
		curr.Next = prev  //修改指针的方向
		prev = curr
		curr = next
	}
	return prev
}

//翻转函数(利用递归)
func reverseList3(head *ListNode) *ListNode {
	//当前函数递归的出口位置
	if head == nil || head.Next == nil {
		return head
	}
	//递归找到链表后面的头节点
	newHead := reverseList3(head.Next)
	head.Next.Next = head //改变链表指针朝向
	head.Next = nil       //将过去的指针报废掉
	return newHead        //返回新的头节点
}
func main() {
	list1 := []int{1, 2, 3, 4, 5}
	ln1 := createList(list1)
	travelList(ln1)
	fmt.Println()
	ln1 = reverseList1(ln1)
	travelList(ln1)
	fmt.Println()
	list2 := []int{1, 2, 3, 4, 5}
	ln2 := createList(list2)
	travelList(ln2)
	fmt.Println()
	ln2 = reverseList2(ln2)
	travelList(ln2)
	fmt.Println()
	list3 := []int{1, 2, 3, 4, 5}
	ln3 := createList(list3)
	travelList(ln3)
	fmt.Println()
	ln3 = reverseList3(ln3)
	travelList(ln3)
}
