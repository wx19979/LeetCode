package main

import "fmt"

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

// 回文链表函数(利用一个线性表存储然后再进行比对)
func isPalindrome1(head *ListNode) bool {
	nums := []int{}   //创建一个空线性表
	for head != nil { //遍历链表
		nums = append(nums, head.Val) //将内容放入其中
		head = head.Next              //移动指针
	}
	n := len(nums)             //获取到数组的长度
	for i := 0; i < n/2; i++ { //逐个的进行比较
		if nums[i] != nums[n-i-1] { //如果出现不符合的情况
			return false //直接返回不是
		}
	}
	return true //整体遍历完没有错误直接返回正确的信息
}

//回文链表函数(利用递归的方式)
func isPalindrome2(head *ListNode) bool {
	//首先创建一个用于对比的前向指针
	frontP := head
	//然后创建一个递归的函数用于产生可以反向比较的指针
	var recursivelyCheck func(*ListNode) bool
	//具体定义函数
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil { //如果当前的不是空的节点
			if !recursivelyCheck(curNode.Next) { //递归判断是否满足该函数的要求(注意这里需要向下递归)
				return false
			}
			if frontP.Val != curNode.Val { //判断两个数值是否相等,出现不想等直接返回错误
				return false
			}
			frontP = frontP.Next //相等的情况就把前向指针继续向后挪
		}
		return true //最终如果都满足情况返回正确信息
	}
	return recursivelyCheck(head) //最终递归调用这个函数
}

//首先实现一个用于反转链表的函数
func reviseList(head *ListNode) *ListNode {
	//首先需要创建一个前向指针还有一个用于操作当前节点的指针
	var prev, cur *ListNode = nil, head
	//开始循环操作
	for cur != nil {
		temp := cur.Next //暂存cur的下一个节点
		cur.Next = prev  //修改cur的指向
		prev = cur       //更新prev
		cur = temp       //更新cur
	}
	return prev //最终返回构造的结果
}

//然后实现一个双指针找到中间节点的函数
func getMiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head                        //首先创建快慢指针
	for fast.Next != nil && fast.Next.Next != nil { //注意这里fast.Next!=nil是用于防止链表只有一个元素的情况
		fast = fast.Next.Next //快指针每次走两格
		slow = slow.Next      //慢指针每次走一格
	}
	return slow //最终返回慢指针
}

//回文链表函数(利用快慢指针的方式)
func isPalindrome3(head *ListNode) bool {
	//首先找到中间的节点
	mid := getMiddleNode(head)
	//默认mid是属于前面的链表的,所以对后面进行反转操作要从mid.Next开始
	secondList := reviseList(mid.Next)
	p1, p2 := head, secondList //创建用于比较的两个链表的指针
	for p2 != nil {            //开始循环比较操作
		if p1.Val != p2.Val { //如果出现不想等的情况
			return false //直接返回错误信息
		}
		p1 = p1.Next //指针向后移动
		p2 = p2.Next
	}
	mid.Next = reviseList(secondList) //最终进行复原操作
	return true                       //返回正确信息
}

func main() {
	list1 := []int{1, 2, 2, 1}
	list_1 := createList(list1)
	travelList(list_1)
	isP1 := isPalindrome1(list_1)
	fmt.Println(isP1)
	list2 := []int{1, 2, 2, 1}
	list_2 := createList(list2)
	travelList(list_2)
	isP2 := isPalindrome2(list_2)
	fmt.Println(isP2)
	list3 := []int{1, 2, 2, 1}
	list_3 := createList(list3)
	travelList(list_3)
	isP3 := isPalindrome3(list_3)
	fmt.Println(isP3)
	travelList(list_3)
}
