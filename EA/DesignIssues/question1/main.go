package main

import (
	"fmt"
	"math/rand"
)

// 定义的类型
type Solution struct {
	data, initState []int
}

// 该对象的构造器
func Constructor(nums []int) Solution {
	//为了和其不共享内存否则会出现错误
	return Solution{nums, append([]int(nil), nums...)}
}

// 该对象的重置函数
func (s *Solution) Reset() []int {
	copy(s.data, s.initState) //注意这里只能用copy否则会出错
	return s.data
}

// 该对象的随机返回数组函数(暴力计算法)
func (s *Solution) Shuffle1() []int {
	//计算数组的长度
	len_d := len(s.data)
	//创建用于存储最终结果的数组
	temp := make([]int, len_d)
	//开始循环操作
	for i := 0; i < len_d; i++ {
		randnum := rand.Intn(len(s.data))                        //随机计算一个值
		temp[i] = s.data[randnum]                                //将随机计算出来的值取出来放到对应位置
		s.data = append(s.data[:randnum], s.data[randnum+1:]...) //将随机计算出来的值从数组中删除掉
	}
	s.data = temp //将最终结果重新赋给数据字段
	return s.data //最终返回数据字段
}

// 第二种方法利用在i到n之间随机抽取一个和第i个位置的元素进行交换操作
func (s *Solution) Shuffle2() []int {
	//获取数据的长度
	len_d := len(s.data)
	//开始循环对其i到len_d之间的随机一个数和第i个数进行交换,重复n次
	for i := 0; i < len_d; i++ {
		//计算从i到n的随机数字
		randnum := i + rand.Intn(len_d-i) //注意这里需要加上i
		//开始交换数据位置
		s.data[i], s.data[randnum] = s.data[randnum], s.data[i]
	}
	return s.data //最终返回结果
}
func main() {
	nums := []int{1, 2, 3}
	obj := Constructor(nums)
	param_1 := obj.Shuffle1()
	fmt.Println(param_1)
	param_2 := obj.Reset()
	fmt.Println(param_2)
	param_3 := obj.Shuffle2()
	fmt.Println(param_3)
}
