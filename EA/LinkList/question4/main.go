package main

import (
	"fmt"
)

// 链表的结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// 首先先定义一个构建链表的函数
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

// 用于遍历的函数
func travelList(L *ListNode) {
	for L.Next != nil {
		fmt.Println(L.Next.Val) //输出结果
		L = L.Next              //继续向下遍历
	}
}

// 合并两个有序链表(迭代法结合尾插法)
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	//首先判断两个链表中有空链表的情况
	if list1 == nil || list2 == nil {
		if list1 == nil { //如果第一个链表是空的就返回第二个
			return list2
		} else { //如果第二个是空的就返回第一个
			return list1
		}
	} else { //两个都不空就需要具体进行操作了
		res := &ListNode{} //首先创建一个用于存储最终结果的一个头指针节点
		tail := res        //创建一个用于进行尾插法的尾指针
		//开始循环比较两个链表的元素
		for list1 != nil && list2 != nil { //注意这里是直接用两个链表的指针进行修改操作
			if list1.Val < list2.Val { //如果出现第一个链表节点值小于第二个链表节点的值
				tail.Next = list1 //利用尾插法把节点插入res的尾部
				tail = list1
				list1 = list1.Next //然后指针继续向后移动
			} else { //否则就把第二个链表的节点插入结果链表中
				tail.Next = list2
				tail = list2
				list2 = list2.Next
			}
		}
		//整个循环完之后可能会有剩下的情况谁剩下了就把剩下的节点插入结果链表的尾部
		if list1 != nil {
			tail.Next = list1
			tail = list1
		} else {
			tail.Next = list2
			tail = list2
		}
		return res.Next //最终返回结果
	}
}

// 递归法
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	//首先写出递归的出口
	if list1 == nil { //如果list1是空的就返回list2
		return list2
	} else if list2 == nil { //如果list2是空的就返回list1
		return list1
	} else if list1.Val < list2.Val { //如果list的节点值小于list的
		list1.Next = mergeTwoLists2(list1.Next, list2) //那么就把list1的节点放到结果链表中
		return list1                                   //并返回list1
	} else { //否则就做相反操作
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}

func main() {
	list1, list2 := []int{4, 2, 1}, []int{4, 3, 1}   //创建两个有序的数组
	Nl1, Nl2 := createList(list1), createList(list2) //根据数组创建链表
	fmt.Println("Nl1:")                              //输出创建好的内容
	travelList(Nl1)
	fmt.Println("Nl2:")
	travelList(Nl2)
	list_1 := mergeTwoLists1(Nl1, Nl2) //执行合并链表的操作
	fmt.Println("Merge1:")
	travelList(list_1)                               //输出合并后的结果
	list3, list4 := []int{4, 2, 1}, []int{4, 3, 1}   //创建两个有序的数组
	Nl3, Nl4 := createList(list3), createList(list4) //根据数组创建链表
	fmt.Println("Nl3:")                              //输出创建好的内容
	travelList(Nl3)
	fmt.Println("Nl4:")
	travelList(Nl4)
	list_2 := mergeTwoLists2(Nl3, Nl4) //执行合并链表的操作
	fmt.Println("Merge2:")
	travelList(list_2) //输出合并后的结果
}
