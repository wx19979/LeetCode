/*
自己定义一个链栈,里面存储指针,数据以及当前节点对应的最小值
,然后直接实现一个栈的正常功能
*/
package main

import (
	"fmt"
	"math"
)

// 重新定义一个新的链栈结构体
type MinStack struct {
	minN int   //全局最小值
	myN  *node //节点指针
}

// 用于进行操作的节点
type node struct {
	data, minNum int   //数据和最小值
	next         *node //链栈的指针
}

// 用于初始化
func Constructor() MinStack {
	return MinStack{minN: math.MaxInt64, //初始化全局最小值
		myN: &node{next: nil}, //初始化节点的指针为空
	}
}

// 入栈操作
func (s *MinStack) Push(val int) {
	s.minN = min(s.minN, val)                                //更新一下全局最小值
	nod := node{minNum: s.minN, data: val, next: s.myN.next} //根据比较结果创建节点(注意这里的指针)
	s.myN.next = &nod                                        //将新节点加入栈中
}

// 出栈操作
func (s *MinStack) Pop() {
	s.myN = s.myN.next     //修改栈顶指针
	if s.myN.next == nil { //如果当前操作后的栈是空栈,需要初始化最小值
		s.minN = math.MaxInt64
	} else { //如果当前操作后不是空栈就更新一下最小值
		s.minN = s.myN.next.minNum
	}
}

// 获取栈顶数据的操作
func (s *MinStack) Top() int {
	return s.myN.next.data
}

// 获取最小值的操作
func (s *MinStack) GetMin() int {
	return s.myN.next.minNum
}

// 用于计算最小值的函数
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	s := Constructor()               //这里直接调用构造函数返回一个预先设定好的对象
	nums := []int{-1, 2, 10, -5, 9}  //创建一个测试数组
	for i := 0; i < len(nums); i++ { //将数据都入栈
		s.Push(nums[i])
	}
	fmt.Println(s.Top())
	min := s.GetMin() //调用获取最小值的函数
	fmt.Println(min)  //输出结果
}
