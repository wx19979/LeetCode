package main

import (
	"fmt"
	"math"
)

// 栈的数据结构(由普通的栈和最小栈组成)
type MinStack struct {
	stack, mstack []int //一个是普通的数据栈,一个是最小值栈
}

// 初始化结构体函数
func Constructor() MinStack {
	return MinStack{
		stack:  []int{},              //普通栈直接初始化为空栈
		mstack: []int{math.MaxInt64}, //最小栈初始要用最大值
	}
}

// 入栈操作
func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)             //将值正常入到普通栈中
	top := s.mstack[len(s.mstack)-1]           //获取到最小栈的栈顶数值
	s.mstack = append(s.mstack, min(val, top)) //比较是当前入栈的值小还是最小栈的值更小,将更小的值入栈
}

// 出栈操作
func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]    //执行普通栈出栈操作
	s.mstack = s.mstack[:len(s.mstack)-1] //执行最小栈的出栈操作
}

// 获取普通栈的栈顶操作
func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1] //返回普通栈栈顶元素的值
}

// 获取最小值就是返回最小栈的栈顶元素
func (s *MinStack) GetMin() int {
	return s.mstack[len(s.mstack)-1] //返回最小值栈栈顶元素的值
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
	min := s.GetMin() //调用获取最小值的函数
	fmt.Println(min)  //输出结果
}
