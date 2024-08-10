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

//计算链表长度的函数
func lenList(head *ListNode) int {
	lenList := 0      //创建一个用于存储长度的变量
	for head != nil { //不断向下进行遍历操作
		lenList++        //进行计数
		head = head.Next //向下遍历
	}
	return lenList //最终返回长度的结果
}

//用于创造环的函数
func createCycle(head *ListNode, pos int) (*ListNode, bool) {
	p1, n := head, lenList(head) //p1是用于存储待扣上环的节点指针,然后获取整个链表的长度
	if n < pos || pos < 0 {      //如果给的位置不合法
		return head, false //直接返回并告知创环失败
	} else {
		for i := 0; i < pos; i++ { //首先遍历到待插入的位置
			p1 = p1.Next //p不断的向下移动
		}
		p2 := head                 //用于遍历到表尾的指针
		for i := 0; i < n-1; i++ { //一直循环到表尾
			p2 = p2.Next
		}
		p2.Next = p1      //将链表表尾指向需要设置的节点处,使之形成环
		return head, true //成功形成环直接返回正确的信息和链表
	}
}

// 判断是否是环形链表(利用两个快慢指针进行操作)
func hasCycle1(head *ListNode) bool {
	//首先排除一些特殊的情况
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head //首先初始化快慢指针
	for fast != slow {            //开始循环移动快慢指针
		if fast == nil || fast.Next == nil { //如果出现快指针遍历到末尾的情况说明不存在环
			return false //返回错误信息
		}
		fast = fast.Next.Next //快指针每次移动两格
		slow = slow.Next      //慢指针每次移动一格
	}
	return true //最终循环能结束说明快慢指针碰在一起了所以直接返回正确信息即可
}

//判断是否是环形链表(利用哈希表进行操作)
func hasCycle2(head *ListNode) bool {
	//首先创建一个哈希表用于存储是否访问过节点的信息
	seen := map[*ListNode]struct{}{}
	//开始循环遍历链表的节点
	for head != nil {
		if _, ok := seen[head]; ok { //如果发现该节点是访问过的节点
			return true //说明有环直接返回信息
		}
		seen[head] = struct{}{} //将当前访问的节点存储到哈希表中
		head = head.Next        //继续向下遍历链表
	}
	return false //最终都遍历完了说明链表不存在环,直接返回信息
}
func main() {
	list1 := []int{3, 2, 0, -4}            //创建一个用于构造链表的数组
	list_1 := createList(list1)            //根据数组创建一个链表
	fmt.Println(hasCycle1(list_1))         //输出链表刚刚创建好的状态应该是无环状态返回false
	list_1, isC1 := createCycle(list_1, 2) //在无环的链表上加环
	if isC1 {                              //如果加环成功
		fmt.Println("Create Success!") //输出信息
		hasC1 := hasCycle1(list_1)     //判断当前是有环
		fmt.Println("hasCycle is:")
		fmt.Println(hasC1) //最终输出信息
	} else { //如果创建失败输出失败的信息
		fmt.Println("Create Fail!")
	}
	list2 := []int{3, 2, 0, -4}
	list_2 := createList(list2)
	fmt.Println(hasCycle1(list_2))
	list_2, isC2 := createCycle(list_2, 2)
	if isC2 {
		fmt.Println("Create Success!")
		hasC2 := hasCycle2(list_2)
		fmt.Println("hasCycle is:")
		fmt.Println(hasC2)
	} else {
		fmt.Println("Create Fail!")
	}
}
