/*
增加一个节点结构体,里面存储好数据和当前节点最小值,
并且需要一个全局最小值与之比较。
*/
package main

import (
	"fmt"
	"math"
)

// 创建一个用于存储最小值的栈
type MinStack struct {
	Mmin  int    //全局最小值
	stack []node //用节点结构体构造栈
}

// 用于构造栈的节点结构体
type node struct {
	data, minNum int //一个是存储数据的字段,一个是存储当前节点的最小值字段
}

// 构造函数
func Constructor() MinStack {
	return MinStack{
		Mmin:  math.MaxInt64, //初始化全局最小值
		stack: []node{},      //栈直接是空的即可
	}
}

// 用于进行入栈操作的函数
func (s *MinStack) Push(val int) {
	s.Mmin = min(s.Mmin, val)              //更新最小的值
	nod := node{data: val, minNum: s.Mmin} //创建待插入节点
	s.stack = append(s.stack, nod)         //将节点插入其中
}

// 用于进行出栈操作
func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1] //直接出栈即可
	//注意这里需要判断当前的栈是否是空栈
	if len(s.stack) == 0 { //如果是空栈的话需要重新初始化一下栈的全局最小值
		s.Mmin = math.MaxInt64
	} else { //如果不是空栈的话将当前栈顶的最小值更新到全局最小值
		s.Mmin = s.GetMin()
	}
}

// 获取栈顶的值
func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].data //直接返回栈顶节点元素的值
}

// 获取最小值
func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].minNum //返回栈顶元素节点的最小值
}

// 计算最小值的函数
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
